package asteroids

import (
	"fmt"

	"github.com/faiface/pixel/text"
	"github.com/tomasaschan/gosteroids/resources/fonts"
	"golang.org/x/image/font"
)

var (
	atariFont  font.Face
	atariAtlas *text.Atlas

	arcadeClassicFont  font.Face
	arcadeClassicAtlas *text.Atlas
)

func init() {
	return
	var err error
	atariFont, err = fonts.AtariFont()
	if err != nil {
		panic(fmt.Errorf("get atari font: %w", err))
	}

	atariAtlas = text.NewAtlas(atariFont, text.ASCII)

	arcadeClassicFont, err = fonts.ArcadeClassic()
	if err != nil {
		panic(fmt.Errorf("get arcade classic font: %w", err))
	}

	arcadeClassicAtlas = text.NewAtlas(arcadeClassicFont, text.ASCII)
}
