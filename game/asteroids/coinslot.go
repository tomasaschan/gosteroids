package asteroids

import (
	"strings"
	"time"

	"github.com/faiface/pixel"
	"github.com/tomasaschan/gosteroids/game/engine"
	"github.com/tomasaschan/gosteroids/game/graphics"
	"golang.org/x/exp/slices"
)

type coinSlot struct {
	quarter bool
}

var _ engine.Drawable = NewCoinSlot()
var _ engine.Ender = NewCoinSlot()
var _ engine.Controlled = NewCoinSlot()

func NewCoinSlot() *coinSlot { return &coinSlot{} }

func (c *coinSlot) Control(pressedKeys []engine.Key, justPressedKeys []engine.Key) {
	if slices.Contains(justPressedKeys, engine.KeyQ) {
		c.quarter = true
	}
}

func (c *coinSlot) EndUpdate(dt time.Duration, objects *engine.GameObjects) {
	if c.quarter {
		objects.Clear()
		objects.Insert(
			NewWaveMaker(),
			NewSaucerMaker(),
			NewShipMaker(),
		)
	}
}

func (*coinSlot) Draw(screen pixel.Target) {
	messages := []string{
		"Insert Coin",
		"[q] - start new game",
	}

	graphics.TextAt(
		screen,
		pixel.V(engine.ScreenSize/2, engine.ScreenSize/4),
		strings.Join(messages, "\n"),
		4,
	)
}
