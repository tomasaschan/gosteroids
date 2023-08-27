# Parallelism Galore

I've had an idea. I want to experiment with goroutines, and how far I can take parallelizing the game engine so that not only do all the independent objects make decisions on their own, draw themselves etc, but they do it _in their own (green) thread_. Go is a language very well suited to this type of idea, so I think I'm in a good place to experiment.

My basic idea is to start by implementing the idea from [the last post][06], where `EndUpdate` returns some type of `Result` instead of actually making changes to the `GameObjects` instance. This would decouple the game objects more from the engine; in particular, it would remove the interaction currently most likely to introduce race conditions.

Once that is done, we can refactor the current method-call mechanisms in the game engine to instead pass messages through channels, allowing all the objects to run individually and only communicate with the engine when they need to.

I don't even know if this will actually improve performance, but it'll be interesting for sure!

[06]: ./06-lifecycle-updates.md