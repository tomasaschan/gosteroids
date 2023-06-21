package asteroids

import (
	"image/color"
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

func (a *Asteroid) InteractWith(other any) {
}

func (a *Asteroid) EndUpdate(dt time.Duration, pressedKeys []ebiten.Key, objects *engine.GameObjects) {
	a.State.Evolve(dt)
}

func (a *Asteroid) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(a.State.P.X), float32(a.State.P.Y), 15, color.White, false)
}
