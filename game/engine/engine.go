package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game interface {
	Name() string
	Slug() any
}

type Drawable interface {
	Draw(*ebiten.Image)
}

type gameEngine struct {
	game Game

	objects GameObjects
}

func (e *gameEngine) Objects() []any {
	return e.objects.objects
}

func (e *gameEngine) Draw(screen *ebiten.Image) {
	for _, o := range e.Objects() {
		if d, ok := o.(Drawable); ok {
			d.Draw(screen)
		}
	}
}

func (*gameEngine) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 1024, 1024
}

func (e *gameEngine) Update() error {
	if e.objects.objects == nil {
		e.objects.Insert(e.game.Slug())
	}

	e.objects.Update(1 / float64(ebiten.TPS()))

	return nil
}

var _ ebiten.Game = &gameEngine{}

func Run(g Game) error {
	ebiten.SetWindowTitle(g.Name())

	return ebiten.RunGame(&gameEngine{game: g})
}
