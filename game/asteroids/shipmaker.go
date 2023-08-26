package asteroids

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/tomasaschan/gosteroids/game/engine"
	"github.com/tomasaschan/gosteroids/game/physics"
)

const ShipEmergenceDelay = 3 * time.Second

type shipMaker struct {
	shipsCreated   int
	lives          int
	seenShip       bool
	emergenceTimer engine.Timer
}

func NewShipMaker() *shipMaker {
	return &shipMaker{
		lives:          3,
		emergenceTimer: engine.NewTimer(ShipEmergenceDelay),
	}
}

var _ engine.Interactor = NewShipMaker()
var _ engine.Ender = NewShipMaker()
var _ engine.Drawable = NewShipMaker()

func (s *shipMaker) InteractWith(other any) {
	if _, ok := other.(*Ship); ok {
		s.seenShip = true
	}
}

func (s *shipMaker) EndUpdate(dt time.Duration, objects *engine.GameObjects) {
	if !s.seenShip {
		s.emergenceTimer.Tick(dt, func() {
			if s.shipsCreated < s.lives {
				s.shipsCreated++
				objects.Insert(NewShip())
			} else {
				fmt.Println("game over")
				objects.Remove(s)
			}
		})
	}
	s.seenShip = false
}

func (s *shipMaker) Draw(target pixel.Target) {
	imd := imdraw.New(nil)

	voffset, hoffset := 130.0, 70.0

	ship := NewShip()
	ship.dropScale = .7

	for i := s.shipsCreated; i < s.lives; i++ {
		ship.State.P = physics.P(hoffset, engine.ScreenSize-voffset)
		ship.Draw(imd)
		hoffset += 40
	}

	imd.Draw(target)
}
