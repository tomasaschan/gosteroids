package asteroids

import (
	"image/color"
	"time"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/tomasaschan/gosteroids/game/engine"
	"github.com/tomasaschan/gosteroids/resources/fonts"
	"golang.org/x/exp/slices"
	"golang.org/x/image/font"
)

var (
	atariFont font.Face
)

func init() {
	var err error
	atariFont, err = fonts.AtariFont()
	if err != nil {
		panic(err)
	}
}

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

		text.Draw(screen, message, atariFont, 1024/2-bbox.Dx()/2, int(800+1.5*float64(line*bbox.Dy())), color.White)
	}
}
