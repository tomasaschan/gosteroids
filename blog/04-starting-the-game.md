# Starting the Game

I think one thing I really like about the distributed design, is that the game engine that orchestrates the advancement of the game, running through the lifecycle methods and interaction calls on each frame, has so little knowledge of the game mechanics. As it stands in [Ron's code as of writing this](https://github.com/RonJeffries/python-asteroids-1/blob/233cb6c75e2262bdf63d5aa9bf2bc8b87d1cfc32/game.py#L34), the game engine is also responsible for accepting coins that reset the game state (i.e. the collection of game objects), but it does so at any point in time - for example, you can start a new game at any point by inserting a quarter (hit `q`), whether it's game over or not.

I'd like to try shifting that responsibility to a (possibly re-usable) game object instead - a `CoinSlot`, if you will - which will accept keyboard input and configure a new game state, but which will only be available to listen for that input when appropriate.

This is also a good opportunity to start figuring out how I want to write tests. Having read the latest few posts in Ron's series, as well as [a commentary article by Rickard Lindberg][high-vs-low] where he explores what he calls "micro" vs "high-level" testing, I think I'm siding slightly with Rickard (insofar as there is any real disagreement at all...); I'd like to avoid writing tests that assume specific call sequences on specific objects, and rather write tests that put carefully constructed sets of objects in the mix, tick time forward, and observe what happens with the total game state.

[high-vs-low]: http://rickardlindberg.me/writing/high-level-low-level-ron-reply/

The `GameObjects` class lends itself quite nicely to this; it currently has the following interface:

```go
// not declared like this, but rather as a struct with these public methods defined...
type GameObjects interface {
    // Remove everything
    Clear()
    // Insert new objects into the mix
    Insert(...any)
    // Remove a specific objects (typically the caller)
    Remove(any)
    // Tick the game forward, through all steps of the call sequence
    Update(dt float64, pressedKeys ...ebiten.Key)
}
```

A test can create a `GameObjects` instance, `Insert` the relevant objects, call `Update` and then run assertions (we might need to add some introspection capabilities for that bit). Passing in the pressed keys to `Update` allows us to test keyboard interactions too, without having to figure out how to interfere with Ebitengine's methods for that!

And, as previously promised, available time to spend on this is limited, and so I won't have time to actually do much of this in this session. There's a `CoinSlot` object now, which does two things:

* It listens for keyboard input and replaces itself with a few asteroids when you hit `q`; still TODO is to refactor this to accept the pressed keys from `Update`
* It renders a message `Insert Coin` with the help text `[q] - start new game` on the screen. Being able to do this was one of the reasons I found this idea compelling (the other being more control over exactly when coins can be inserted).

Until next time!