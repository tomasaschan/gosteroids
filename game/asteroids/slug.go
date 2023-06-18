package asteroids

import (
	"github.com/tomasaschan/gosteroids/game/engine"
	"github.com/tomasaschan/gosteroids/game/physics"
)

type slug struct{}

var _ engine.Ender = &slug{}

func (s *slug) EndUpdate(dt float64, objects *engine.GameObjects) {
	objects.Remove(s)

	objects.Insert(
		&Asteroid{State: physics.State{X: 400, Y: 500, Vx: 100, Vy: 20, Theta: 0, Vtheta: 10}},
		&Asteroid{State: physics.State{X: 300, Y: 800, Vx: -50, Vy: 50, Theta: 0, Vtheta: 10}},
		&Asteroid{State: physics.State{X: 600, Y: 100, Vx: 20, Vy: -120, Theta: 0, Vtheta: 10}},
	)
}
