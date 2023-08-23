package main

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/tomasaschan/gosteroids/resources/fonts"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Gosteroids",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	canvas := win.Canvas()
	// imd := imdraw.New(canvas)
	arcadeClassic, err := fonts.ArcadeClassic()
	if err != nil {
		panic(err)
	}
	atlas := text.NewAtlas(arcadeClassic, text.ASCII)

	targetFrameDuration := 1 * time.Second / 60
	currentFrameRate := 60
	frames := 0

	last := time.Now()
	lastPrint := time.Now()
	var dtPrint time.Duration

	for !win.Closed() {
		frames++
		dt := time.Since(last)

		if dt < targetFrameDuration {
			time.Sleep(targetFrameDuration - dt)
		}
		last = time.Now()

		canvas.Clear(colornames.Midnightblue)
		txt := text.New(pixel.V(50, 100), atlas)
		txt.Color = colornames.White
		fmt.Fprintln(txt, "hello, pixel")

		if time.Since(lastPrint) > 1*time.Second {
			currentFrameRate, frames = frames, 0
			dtPrint = time.Since(lastPrint)
			lastPrint = time.Now()
		}

		fmt.Fprintf(txt, "current framerate: %d / %s", currentFrameRate, dtPrint)

		txt.Draw(canvas, pixel.IM)

		win.Update()
	}

}

func main() {
	pixelgl.Run(run)

	fmt.Println("end of program!")
}
