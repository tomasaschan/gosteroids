package asteroids

import (
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/tomasaschan/gosteroids/game/engine"
	"github.com/tomasaschan/gosteroids/game/physics"
	"golang.org/x/exp/slices"
)

const shipRadius = 35

var ShipPoints = []pixel.Vec{
	pixel.V(-15, -10),
	pixel.V(-15, 10),
	pixel.V(-25, 20),
	pixel.V(35, 0),
	pixel.V(-25, -20),
	pixel.V(-15, -10),
}

type Ship struct {
	State     physics.State
	boosting  bool
	dropScale float64
	colliding bool
}

func NewShip() *Ship {
	return &Ship{
		State:     physics.NewStationary(engine.ScreenSize/2, engine.ScreenSize/2, math.Pi/2),
		dropScale: 50,
	}
}

var _ engine.Interactor = NewShip()
var _ engine.Controlled = NewShip()
var _ engine.Ender = NewShip()
var _ engine.Drawable = NewShip()

func (s *Ship) InteractWith(other any) {
	if asteroid, ok := other.(*asteroid); ok {
		if asteroid.CollidingWith(s.State.P, shipRadius) {
			s.colliding = true
		}
	}
}

func (s *Ship) Control(pressedKeys []engine.Key, justPressedKeys []engine.Key) {
	s.boosting = false
	s.State.Vtheta = 0

	if s.dropScale > 1 {
		// can't control the ship before it's dropped in fully
		return
	}

	if slices.Contains(pressedKeys, engine.KeyLeftArrow) {
		s.State.Vtheta += math.Pi
	}
	if slices.Contains(pressedKeys, engine.KeyRightArrow) {
		s.State.Vtheta -= math.Pi
	}
	if slices.Contains(pressedKeys, engine.KeyUpArrow) {
		s.State.V = s.State.V.Add(physics.P(1, 0).Rotate(s.State.Theta).Scale(3))
		s.boosting = true
	}
}

func (s *Ship) EndUpdate(dt time.Duration, objects *engine.GameObjects) {
	if s.dropScale > 1 {
		s.dropScale -= dt.Seconds() * 100
	}
	if s.dropScale < 1 {
		s.dropScale = 1
	}

	if s.colliding {
		objects.Remove(s)
		objects.Insert(s.Eplode()...)
		return
	}

	s.State.Evolve(dt)
	s.colliding = false
}

func (s *Ship) Draw(target pixel.Target) {
	imd := imdraw.New(nil)

	imd.SetMatrix(pixel.IM.
		Rotated(pixel.ZV, s.State.Theta).
		Scaled(pixel.ZV, s.dropScale).
		Moved(pixel.Vec(s.State.P)))
	imd.Push(ShipPoints...)
	imd.Polygon(4)

	if s.boosting {
		imd.Push(pixel.V(-15, -10), pixel.V(-35, 0), pixel.V(-15, 10))
		imd.Polygon(4)
	}

	imd.Draw(target)
}

func (s *Ship) Eplode() (result []any) {
	for i := 0; i < len(ShipPoints)-1; i++ {
		result = append(result, NewLineFragment(s.State, ShipPoints[i], ShipPoints[i+1]))
	}

	return
}
