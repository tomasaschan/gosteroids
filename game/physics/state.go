package physics

import "time"

type State struct {
	P             Point
	V             Point
	Theta, Vtheta float64
}

func evolve(x, v float64, dt time.Duration) float64 {
	x += v * dt.Seconds()
	for x < 0 {
		x += 1024
	}
	for 1024 < x {
		x -= 1024
	}

	return x
}

func (s *State) Evolve(dt time.Duration) {
	s.P.X = evolve(s.P.X, s.V.X, dt)
	s.P.Y = evolve(s.P.Y, s.V.Y, dt)

	s.Theta += s.Vtheta * dt.Seconds()
}
