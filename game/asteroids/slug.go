package asteroids

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tomasaschan/gosteroids/game/engine"
)

type slug struct{}

var _ engine.Ender = &slug{}

func (s *slug) EndUpdate(dt time.Duration, keys []ebiten.Key, objects *engine.GameObjects) {
	objects.Clear()

	objects.Insert(NewCoinSlot(), NewWaveMaker())
}
