# A blank canvas

Since I've never written a game before, in Golang or otherwise, I started out by searching <sup id="searching-back">[1](#searching)</sup> for game engine libraries. On my own, I foumd [pixel](https://github.com/faiface/pixel) which looks promising but complicated - but several others recommended [Ebitengine][ebitengine], and its tagline "A dead simple 2D game engine for Go" sounds like _exactly_ what I'm looking for.

To get started, I built a simple animation of a ball that flies around the window. When it reaches the edge, half the time it bounces and half the time it wraps.

A few things I learned today:

* How to set up a game loop and do some very stupid physics simulation in it
* That Ebitengine automatically scales my window for me, so I can decide the precision I want when drawing stuff (by setting my virtual screen size) and go from there

Next steps is probably to try to create the same thing, but start approaching the target architecture, with mutually interacting game objects that create the game itself as an emergent behavior, and a wrapping game engine that just gives all the game objects a chance to do their thing. We'll see how long it takes before I attempt that! 😅

[ebitengine]: https://ebitengine.org/

--- 

<sup id="searching">1)</sup> And by that, I mean both using [internet search engines](https://duckduckgo.com) and [flesh search engines](https://hachyderm.io/@tomasaschan/110550163007847871) [↩](#searching-back)