# A New Engine

It's been a while! The last update to this project was back in early July, almost two months ago. Can't say I didn't warn you updates would be irregular, though...

There are two main reasons for this:

1. It's been summer vacation time, and I've been spending time with my family, reading books, etc, rather than coding on nonsense projects.

2. When I broke off, I had a weird performance issue that I didn't feel like digging into. My guess is that it was somehow related to the combination of OpenGL, Arch Linux, WSL and Windows, where the chain was severed at some link the game engine I was using (Ebiten) to Windows. This resulted in a deeply unplayable handful of frames per second.

Over the past few days I've been ripping out Ebiten and replacing it with [Pixel] - potentially, unknowingly changing some other links in that chain as well - and I'm now able to achieve near-enough 60 FPS with mostly the same code for the actual game. We're back in business! Or, as it were, pleasure, since I don't expect anyone will ever pay me for this...

I also made some minor changes to the game engine design; I'll write something up on that in a separate post another time.

With this, I think I'm ready to start actually implementing Asteroids!

[Ebiten]: https://ebitengine.org/
[Pixel]: https://github.com/faiface/pixel