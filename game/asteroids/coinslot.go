package asteroids

import (
	"image/color"
	"time"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/tomasaschan/gosteroids/game/engine"
	"golang.org/x/exp/slices"
)

type coinSlot struct{}

var _ engine.Drawable = coinSlot{}
var _ engine.Ender = coinSlot{}

func NewCoinSlot() coinSlot { return coinSlot{} }

func (coinSlot) EndUpdate(dt time.Duration, pressedKeys []ebiten.Key, objects *engine.GameObjects) {
	if slices.Contains(pressedKeys, ebiten.KeyQ) {
		objects.Clear()
		objects.Insert()
	}
}

func (coinSlot) Draw(screen *ebiten.Image) {
	messages := []string{
		"Insert Coin",
		"[q] - start new game",
	}

	for line, message := range messages {
		bbox := text.BoundString(atariFont, message)

		text.Draw(screen, message, atariFont, engine.ScreenSize/2-bbox.Dx()/2, int(engine.ScreenSize*2/3+1.5*float64(line*bbox.Dy())), color.White)
	}
}
