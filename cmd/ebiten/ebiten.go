package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tomasaschan/gosteroids/game/hello"
)

func main() {
	ebiten.SetWindowTitle("Hello, world!")
	ebiten.SetFullscreen(true)

	if err := ebiten.RunGameWithOptions(hello.NewGame(), &ebiten.RunGameOptions{}); err != nil {
		log.Fatal(err)
	}
}
