package asteroids

import (
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
