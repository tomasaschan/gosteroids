package engine

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const ScreenSize = 1024

type Game interface {
	Name() string
	Init() any
	BackgroundColor() color.Color
}

type Drawable interface {
	Draw(*ebiten.Image)
}

type gameEngine struct {
	game    Game
	keys    []ebiten.Key
	objects GameObjects
}

func (e *gameEngine) Objects() []any {
	return e.objects.objects
}

func (e *gameEngine) Draw(screen *ebiten.Image) {
	screen.Fill(e.game.BackgroundColor())
	for _, o := range e.Objects() {
		if d, ok := o.(Drawable); ok {
			d.Draw(screen)
		}
	}
}

func (*gameEngine) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return ScreenSize, ScreenSize
}

func (e *gameEngine) Update() error {
	if e.objects.objects == nil {
		e.objects.Insert(e.game.Init())
	}

	e.keys = inpututil.AppendPressedKeys(e.keys[:0])

	e.objects.Update(time.Duration(1000/float64(ebiten.TPS()))*time.Millisecond, e.keys...)

	return nil
}

var _ ebiten.Game = &gameEngine{}

func Run(g Game) error {
	ebiten.SetWindowTitle(g.Name())

	return ebiten.RunGame(&gameEngine{game: g})
}
