package asteroids

import (
	"github.com/tomasaschan/gosteroids/game/engine"
)

type slug struct{}

var _ engine.Ender = &slug{}

func (s *slug) EndUpdate(dt float64, objects *engine.GameObjects) {
	objects.Remove(s)

	objects.Insert(
		&coinSlot{},
	)
}
