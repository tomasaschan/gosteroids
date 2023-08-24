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

type ship struct {
	State     physics.State
	boosting  bool
	dropScale float64
}

func NewShip() *ship {
	return &ship{
		State:     physics.NewStationary(engine.ScreenSize/2, engine.ScreenSize/2, math.Pi/2),
		dropScale: 50,
	}
}

var _ engine.Controlled = NewShip()
var _ engine.Ender = NewShip()
var _ engine.Drawable = NewShip()

func (s *ship) Control(pressedKeys []engine.Key, justPressedKeys []engine.Key) {
	s.boosting = false
	s.State.Vtheta = 0

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

func (s *ship) EndUpdate(dt time.Duration, objects *engine.GameObjects) {
	if s.dropScale > 1 {
		s.dropScale -= dt.Seconds() * 100
		return
	} else if s.dropScale < 1 {
		s.dropScale = 1
	}

	s.State.Evolve(dt)
}

func (s *ship) Draw(target pixel.Target) {
	imd := imdraw.New(nil)
	imd.SetMatrix(pixel.IM.Rotated(pixel.ZV, s.State.Theta).Scaled(pixel.ZV, 5*s.dropScale).Moved(pixel.Vec(s.State.P)))
	imd.Push(pixel.V(-3, -2), pixel.V(-3, 2), pixel.V(-5, 4), pixel.V(7, 0), pixel.V(-5, -4), pixel.V(-3, -2))
	imd.Polygon(1)

	if s.boosting && s.dropScale == 1 {
		imd.Push(pixel.V(-3, -2), pixel.V(-7, 0), pixel.V(-3, 2))
		imd.Polygon(1)
	}

	imd.Draw(target)
}
