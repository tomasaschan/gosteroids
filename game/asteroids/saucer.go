package asteroids

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/tomasaschan/gosteroids/game/engine"
	"github.com/tomasaschan/gosteroids/game/graphics"
	"github.com/tomasaschan/gosteroids/game/physics"
)

type saucer struct {
	State physics.State
}

func NewSaucer(p physics.Point, v physics.Point) *saucer {
	return &saucer{State: physics.State{P: p, V: v}}
}

func (s *saucer) Draw(target pixel.Target) {
	graphics.TextAt(target, pixel.Vec(s.State.P), "<saucer>", 2)
}

func (s *saucer) EndUpdate(dt time.Duration, objects *engine.GameObjects) {
	s.State.Evolve(dt)
}

var _ engine.Drawable = NewSaucer(physics.Zero, physics.Zero)
var _ engine.Ender = NewSaucer(physics.Zero, physics.Zero)
