package asteroids

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tomasaschan/gosteroids/game/engine"
	"github.com/tomasaschan/gosteroids/game/physics"
)

type Asteroid struct {
	State physics.State
}

var _ engine.Interactor = &Asteroid{}
var _ engine.Ender = &Asteroid{}
var _ engine.Drawable = &Asteroid{}

func NewAsteroid() *Asteroid {
	p := physics.Point{X: 0, Y: rand.Float64() * engine.ScreenSize}
	v := physics.Point{X: 100, Y: 0}.Rotate(rand.Float64() * 360)

	a := &Asteroid{State: physics.State{P: p, V: v, Theta: 0, Vtheta: 10}}
	// ensure asteroid appears on the correct side of the screen on first draw
	a.State.Evolve(1 * time.Millisecond)

	return a
}

func (a *Asteroid) InteractWith(other any) {
}

func (a *Asteroid) EndUpdate(dt time.Duration, pressedKeys []ebiten.Key, objects *engine.GameObjects) {
	a.State.Evolve(dt)
}

func (a *Asteroid) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(a.State.P.X), float32(a.State.P.Y), 15, color.White, false)
}
