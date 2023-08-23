package physics

import (
	"time"

	"github.com/tomasaschan/gosteroids/game/engine"
)

type State struct {
	P             Point
	V             Point
	Theta, Vtheta float64
}

func S0(x, y, theta float64) State {
	return State{
		P:     P(x, y),
		Theta: theta,
	}
}

func evolve(x, v float64, dt time.Duration) float64 {
	x += v * dt.Seconds()
	for x < 0 {
		x += engine.ScreenSize
	}
	for engine.ScreenSize < x {
		x -= engine.ScreenSize
	}

	return x
}

func (s *State) Evolve(dt time.Duration) {
	s.P.X = evolve(s.P.X, s.V.X, dt)
	s.P.Y = evolve(s.P.Y, s.V.Y, dt)

	s.Theta += s.Vtheta * dt.Seconds()
}
