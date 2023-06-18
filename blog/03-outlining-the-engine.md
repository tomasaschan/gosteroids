# Outlining the game engine

I think we're getting somewhere!

In Ron's articles, a lot of the meat is in the exploration of how to radically change the architecture in small, safe steps. While that is very interesting in itself, I'm doing this more to explore the design decisions themselves than to explore the steps to get there - and I'm also cheating a bit, since I've been thinking about his design for a while (specifically, for as long as I've been reading his posts). I also don't have an existing codebase I'm trying to move to the new architecture - I just want to implement it from scratch and see what it feels like.

In this iteration, I've started iterating on the game-agnostic engine, and on some very crude building blocks for an Asteroids game.

We have the following components in place for the engine itself:

* A `GameEngine` which implements the `ebiten.Game` interface, and acts as a generic bridge to our idea of a distributed game.
* `GameObjects`, a "smart" collection that keeps track of game objects and orchestrates their interactions. The lifecycle of a frame is currently three simple methods:
  * `object.BeginUpdate(dt float64)`, sent to all objects individually at the start of a `Update` (in the ebiten sense). This is where objects can reset any internal state for keeping track of interactions etc, and where they should move.
  * `object.InteractWith(other)`, sent to all object pairs (every `object` will see every `other`, but not itself)
  * `object.EndTick(dt float64, objects *GameObjects)`, sent to all objects individually after all interactions have been handled. This is where an object takes action based on interactions (e.g. asteroids remove themselves after colliding with something, maybe adding smaller ones in their place) or timers (e.g. missiles time out).

This lifecycle is slightly simplified from Ron's, and maybe less expressive, but I think it'll suffice.

We also have the embryo to the Asteroids game:

* A `Slug` that will be responsible for starting "attract mode"
* An `Asteroid` that can move around on the screen and draw itself (as a circle, for now)

I've made a few decisions that differ from Rob's, at least partly because I'm in a different language with different features:

* He's trying to do all the necessary type checking when interacting by double dispatch. I'm going to do it instead with safe casts - `if a, ok := o.(Asteroid); ok { /* a is now an Asteroid */ }`, removing the need for all objects to implement `interactWithX` for all other object types `X`.

* For optional lifecycle methods, I use one-method interfaces and safe casts to see if the methods should be called, rather than providing default noop implementations. I can still compile-time check that they are properly implemented, with declarations such as `var _ engine.Interactor = &Asteroid{}`, which will yield a compile error if the lifecycle method does not exist with the right signature.

* I also plan to have no control mechanisms in the engine itself - all user interactions, including inserting quarters etc, will be handled by game objects in the mix, rather than by the central engine. This allows me to e.g. control when a quarter can be inserted, by having an object that accepts quarters that can be present or not as needed.

You'll note that I'm not very careful about committing stuff in small steps - or writing tests. This might change in the future, but given my constraints of limited available hack time, I prefer to just hack away and commit when I feel "done" with the session.

Next steps are probably to create a WaveMaker for asteroids, and to render them a bit nicer.