package engine

import "github.com/faiface/pixel/pixelgl"

// Key abstracts away the constants used for keyboard input in various game engines
// so we don't have to change code everywhere if we switch out engines
type Key int

// If I were diligent, these constants should bring over all keys on the keyboard, but that's tedious.
// Instead, we just list the ones we actually make use of.
// This might be an opportunity to learn how go source generators work.
const (
	KeyQ          = Key(pixelgl.KeyQ)
	KeyLeftArrow  = Key(pixelgl.KeyLeft)
	KeyRightArrow = Key(pixelgl.KeyRight)
	KeyUpArrow    = Key(pixelgl.KeyUp)
)

var keys = []Key{
	KeyQ,
	KeyLeftArrow,
	KeyRightArrow,
	KeyUpArrow,
}

func KeyboardInput(w *pixelgl.Window) (pressedKeys []Key, justPressedKeys []Key) {
	for _, key := range keys {
		if w.Pressed(pixelgl.Button(key)) {
			pressedKeys = append(pressedKeys, key)
		}
		if w.JustPressed(pixelgl.Button(key)) {
			justPressedKeys = append(justPressedKeys, key)
		}
	}

	return
}
