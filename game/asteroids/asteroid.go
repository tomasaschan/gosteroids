package asteroids

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tomasaschan/gosteroids/game/engine"
	"github.com/tomasaschan/gosteroids/game/physics"
)

type Asteroid struct {
	State physics.State
}

var _ engine.Interactor = &Asteroid{}
var _ engine.Beginner = &Asteroid{}
var _ engine.Drawable = &Asteroid{}

func (a *Asteroid) BeginUpdate(dt float64) {
	a.State.Evolve(dt)
}

func (a *Asteroid) InteractWith(other any) {
}

func (a *Asteroid) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(a.State.X), float32(a.State.Y), 15, color.White, false)
}
