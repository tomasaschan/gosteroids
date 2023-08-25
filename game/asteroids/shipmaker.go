package asteroids

import (
	"time"

	"github.com/tomasaschan/gosteroids/game/engine"
)

const ShipEmergenceDelay = 3 * time.Second

type shipMaker struct {
	livesRemaining int
	seenShip       bool
	emergenceTimer engine.Timer
}

func NewShipMaker() *shipMaker {
	return &shipMaker{
		livesRemaining: 3,
		emergenceTimer: engine.NewTimer(ShipEmergenceDelay),
	}
}

var _ engine.Interactor = NewShipMaker()
var _ engine.Ender = NewShipMaker()

func (s *shipMaker) InteractWith(other any) {
	if _, ok := other.(*Ship); ok {
		s.seenShip = true
	}
}

func (s *shipMaker) EndUpdate(dt time.Duration, objects *engine.GameObjects) {
	if !s.seenShip {
		s.emergenceTimer.Tick(dt, func() {
			objects.Insert(NewShip())
		})
	}
	s.seenShip = false
}
