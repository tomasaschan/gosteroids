package asteroids

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/tomasaschan/gosteroids/game/engine"
	"github.com/tomasaschan/gosteroids/game/graphics"
)

type message struct {
	text      string
	longevity time.Duration
	age       time.Duration
}

func NewMessage(text string, longevity time.Duration) *message {
	return &message{
		text:      text,
		longevity: longevity,
	}
}

var _ engine.Actor = NewMessage("", 0)
var _ engine.Drawable = NewMessage("", 0)

func (m *message) Act(dt time.Duration) engine.Result {
	if m.age > m.longevity {
		return engine.Result{RemoveSelf: true}
	}

	m.age += dt
	return engine.Result{}
}

func (m *message) Draw(t pixel.Target) {
	graphics.TextAt(t, pixel.V(engine.ScreenSize/2, engine.ScreenSize/2), m.text, 5)
}
