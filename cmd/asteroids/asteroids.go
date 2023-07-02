package main

import (
	"github.com/tomasaschan/gosteroids/game/asteroids"
	"github.com/tomasaschan/gosteroids/game/engine"

	_ "github.com/silbinarywolf/preferdiscretegpu"
)

func main() {
	engine.Run(asteroids.NewGame())
}
