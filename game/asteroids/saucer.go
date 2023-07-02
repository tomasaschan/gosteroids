package asteroids

import (
	"image/color"
	"time"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/tomasaschan/gosteroids/game/engine"
	"github.com/tomasaschan/gosteroids/game/physics"
)

type saucer struct {
	State physics.State
}

func NewSaucer(p physics.Point, v physics.Point) *saucer {
	return &saucer{State: physics.State{P: p, V: v}}
}

func (s *saucer) Draw(screen *ebiten.Image) {
	text.Draw(screen, "<saucer>", atariFont, int(s.State.P.X), int(s.State.P.Y), color.White)
}

func (s *saucer) EndUpdate(dt time.Duration, pressedKeys []ebiten.Key, objects *engine.GameObjects) {
	s.State.Evolve(dt)
}

var _ engine.Drawable = &saucer{}
var _ engine.Ender = &saucer{}
