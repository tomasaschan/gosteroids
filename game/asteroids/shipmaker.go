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
	shipsCreated int
	lives        int
	seenShip     bool
	waited       time.Duration
}

func NewShipMaker() *shipMaker {
	return &shipMaker{lives: 3}
}

var _ engine.Interactor = NewShipMaker()
var _ engine.Actor = NewShipMaker()
var _ engine.Drawable = NewShipMaker()

func (s *shipMaker) InteractWith(other any) {
	if _, ok := other.(*Ship); ok {
		s.seenShip = true
	}
}

func (s *shipMaker) Act(dt time.Duration) (result engine.Result) {
	if !s.seenShip {
		s.waited += dt
	}

	if s.waited > ShipEmergenceDelay {
		s.waited = 0
		if s.shipsCreated < s.lives {
			s.shipsCreated++
			result.NewObjects = append(result.NewObjects, NewShip())
		} else {
			fmt.Println("game over")
			result.RemoveSelf = true
		}
	}

	s.seenShip = false

	return
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
