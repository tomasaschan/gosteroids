package asteroids

import (
	"image/color"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/tomasaschan/gosteroids/game/engine"
	"github.com/tomasaschan/gosteroids/game/physics"
	"github.com/tomasaschan/gosteroids/resources/fonts"
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

var _ engine.Drawable = &coinSlot{}
var _ engine.Ender = &coinSlot{}

func (*coinSlot) EndUpdate(dt float64, objects *engine.GameObjects) {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		objects.Clear()
		objects.Insert(
			&Asteroid{State: physics.State{X: 400, Y: 500, Vx: 100, Vy: 20, Theta: 0, Vtheta: 10}},
			&Asteroid{State: physics.State{X: 300, Y: 800, Vx: -50, Vy: 50, Theta: 0, Vtheta: 10}},
			&Asteroid{State: physics.State{X: 600, Y: 100, Vx: 20, Vy: -120, Theta: 0, Vtheta: 10}},
		)
	}
}

func (*coinSlot) Draw(screen *ebiten.Image) {
	messages := []string{
		"Insert Coin",
		"[q] - start new game",
	}

	for line, message := range messages {
		bbox := text.BoundString(atariFont, message)

		text.Draw(screen, message, atariFont, 1024/2-bbox.Dx()/2, int(800+1.5*float64(line*bbox.Dy())), color.White)
	}
}
