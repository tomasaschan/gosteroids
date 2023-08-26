package asteroids

import "github.com/tomasaschan/gosteroids/game/physics"

func AreColliding(a, b physics.State, totalRadius float64) bool {
	return a.P.DistanceTo(b.P) < totalRadius
}
