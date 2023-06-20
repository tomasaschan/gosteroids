package asteroids

import (
	"image/color"

	"github.com/tomasaschan/gosteroids/game/engine"
)

type asteroids struct{}

var _ engine.Game = &asteroids{}

func (a *asteroids) Name() string { return "Asteroids" }

func (a *asteroids) Init() any {
	return &slug{}
}

func NewGame() *asteroids {
	return &asteroids{}
}

func (*asteroids) BackgroundColor() color.Color {
	return color.RGBA{R: 0x00, G: 0x0a, B: 0x30}
}
