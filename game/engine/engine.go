package engine

import (
	"image/color"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const ScreenSize = 1024

type Game interface {
	Name() string
	Init() []any
	BackgroundColor() color.Color
}

type Drawable interface {
	Draw(pixel.Target)
}

type gameEngine struct {
	window *pixelgl.Window

	game    Game
	keys    []Key
	objects GameObjects
}

func (*gameEngine) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return ScreenSize, ScreenSize
}

func (e *gameEngine) Update(dt time.Duration, pressedKeys []Key, justPressedKeys []Key) {
	if e.objects.objects == nil {
		e.objects.Insert(e.game.Init()...)
	}

	e.objects.Update(dt, pressedKeys, justPressedKeys)
}

func (e *gameEngine) RunFrame(dt time.Duration) {
	pressed, justPressed := KeyboardInput(e.window)
	e.window.Canvas().Clear(colornames.Midnightblue)

	e.Update(dt, pressed, justPressed)
	e.objects.Draw(e.window.Canvas())

	e.window.Update()
}

func (e *gameEngine) Running() bool {
	return !e.window.Closed()
}

func NewEngine(g Game) *gameEngine {
	cfg := pixelgl.WindowConfig{
		Title:  g.Name(),
		Bounds: pixel.R(0, 0, ScreenSize, ScreenSize),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	return &gameEngine{game: g, window: win}
}

func Run(g Game) {
	pixelgl.Run(func() {
		engine := NewEngine(g)

		last := time.Now()
		for engine.Running() {
			dt := time.Since(last)
			last = time.Now()
			engine.RunFrame(dt)
		}
	})
}
