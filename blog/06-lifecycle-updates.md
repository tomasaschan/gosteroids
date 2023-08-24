# Lifecycle Updates

## I'm Way Ahead of You

Yesterday's post about the switch of underlying game engine library hinted at some changes to the lifecycle methods; I thought I'd write those up before I go ahead and make even more changes I have in mind.

Back in [Outlining the Engine][03], I said we'd have the following lifecycle methods (already updated to reflect that, in [Starting the Game][04], I added a parameter for keyboard input to `EndTick`, to allow game objects to accept keyboard input without making it too hard to write tests):

* `object.BeginUpdate(dt float64)`
* `object.InteractWith(other)`
* `object.EndTick(dt float64, pressedKeys []ebiten.Key, objects *GameObjects)`

I then had to yet another parameter to distinguish between keys that are currently being held in, and keys that were _just_ pressed down. For example, when moving the ship one should be able to just hold down the arrow keys and accellerate, while when firing the ship's cannon, we should only react _once_ and then one should have to release the trigger before one can fire again.

This quickly became pretty annoying, since most game objects don't listen to keyboard input at all and just ignored those parameters - an interface split begging to be born! I ended up with the following lifecycle interfaces:

```go
type Beginner interface {
    BeginUpdate()
}
type Interactor interface { 
    InteractWith(other any)
}
type Controlled interface {
	Control(pressedKeys []Key, justPressedKeys []Key)
}
type Ender interface {
	EndUpdate(dt time.Duration, objects *GameObjects)
}
```

(I'm still not super happy about the names; they might change in the future...)

## Where Do We Go From Here?

I have a couple of ideas for furhter refinement:

First, I'm tempted to either merge `Beginner` and `Interactor`, or remove `Beginner` entirely. Interactors should basically always reset some state before starting a new round of interactions, and it's just a matter of style how we do that - and we're already at a point where no actual action is taken until the `EndUpdate` call anyway.

* With both `Begin` and `InteractWith(other)` calls, then all interactors should implement `Begin`. _Maybe_ there are cases for an empty method, but it should be very rare and cheap in those cases.

* With only `InteractWith(other)`, the contract would instead be that the state should be left "ready for a new cycle" at the end of `EndUpdate` (or on creation). This feels "cleaner", but is also a bit harder to explain - and possibly a bit harder to get right.

I'm on the fence, but leaning toward the second option - if for no other reason, then because I find it harder to know what I think about it before I see it in code.

Secondly, I would like to see if I can remove the `*GameObjects` parameter to `EndUpdate` - while handy, it does give each object a lot of power to do stuff it shouldn't (e.g. reference or even remove other objects at will). I'd like to experiment with returning some kind of result struct that indicates actions the `GameObjects` collection can take, including

* removing the currently acting game object
* clearing out all game objects (useful for e.g. `CoinSlot` when starting a new game)
* adding new game objects

If I make both these changes (I'm just leaving them on the table for now), we'd end up with interfaces maybe resembling this:

```go
type Result struct {
    RemoveSelf  bool
    RemoveAll   bool
    NewObjects  []any
}

type Interactor interface { 
    InteractWith(other any)
}
type Controlled interface {
	Control(pressedKeys []Key, justPressedKeys []Key)
}
type Actor interface {
	Act(dt time.Duration) Result
}
```

[03]: ./03-outlining-the-engine.md
[04]: ./04-starting-the-game.md