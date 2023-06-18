package main

import (
	"github.com/tomasaschan/gosteroids/game/asteroids"
	"github.com/tomasaschan/gosteroids/game/engine"
)

func main() {
	engine.Run(asteroids.NewGame())
}
