package asteroids

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/tomasaschan/gosteroids/game/engine"
	"github.com/tomasaschan/gosteroids/game/physics"
)

type asteroid struct {
	State physics.State
}

var _ engine.Interactor = NewAsteroid()
var _ engine.Ender = NewAsteroid()
var _ engine.Drawable = NewAsteroid()

func NewAsteroid() *asteroid {
	p := physics.Point{X: 0, Y: rand.Float64() * engine.ScreenSize}
	v := physics.Point{X: 100, Y: 0}.Rotate(rand.Float64() * 360)

	a := &asteroid{State: physics.State{P: p, V: v, Theta: 0, Vtheta: 10}}
	// ensure asteroid appears on the correct side of the screen on first draw
	a.State.Evolve(1 * time.Millisecond)

	return a
}

func (a *asteroid) InteractWith(other any) {
}

func (a *asteroid) EndUpdate(dt time.Duration, objects *engine.GameObjects) {
	a.State.Evolve(dt)
}

func (a *asteroid) Draw(screen pixel.Target) {
	imd := imdraw.New(nil)

	imd.Color = color.White
	imd.Push(pixel.Vec(a.State.P))
	imd.Circle(15, 0)

	imd.Draw(screen)
}
