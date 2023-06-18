package physics

type State struct {
	X, Y,
	Vx, Vy,
	Theta, Vtheta float64
}

func evolve(x, v, dt float64) float64 {
	x += v * dt
	for x < 0 {
		x += 1024
	}
	for 1024 < x {
		x -= 1024
	}

	return x
}

func (s *State) Evolve(dt float64) {
	s.X = evolve(s.X, s.Vx, dt)
	s.Y = evolve(s.Y, s.Vy, dt)

	s.Theta += s.Vtheta * dt
}
