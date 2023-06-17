package hello

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	h = 640
	w = 480
)

type Game struct {
	x, y, vx, vy float32

	bounces, wraps int
}

func NewGame() *Game {
	return &Game{
		vx: 150,
		vy: 150,
	}
}

func mod(x float32, s int) float32 {
	return float32(math.Mod(float64(x), float64(s)))
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(
		screen,
		mod(g.x, w),
		mod(g.y, h),
		10,
		color.White,
		true,
	)

	fps := fmt.Sprint("frame rate ", ebiten.ActualFPS(), ebiten.TPS())
	if g.bounces+g.wraps > 0 {
		ebitenutil.DebugPrint(screen, fmt.Sprint(
			fps, "\n",
			"bounces ", g.bounces, 100.*g.bounces/(g.bounces+g.wraps), "% | ", 100.*g.wraps/(g.bounces+g.wraps), "% ", g.wraps, " wraps"))
	} else {
		ebitenutil.DebugPrint(screen, fps)
	}

}

func (*Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return w, h
}

func (g *Game) Update() error {
	dt := float32(1.0 / ebiten.DefaultTPS)

	maxx, maxy := g.Layout(0, 0)

	update := func(p, v *float32, mx int) {
		*p += *v * dt
		if *p < 0 || *p > float32(mx) {
			if rand.Float32() < .5 {
				// bounce half the time
				g.bounces++
				*v = -*v
				*p += 2 * *v * dt
			} else {
				// wrap the other half
				g.wraps++
				*p = mod(*p, mx)
			}
		}
	}

	update(&g.x, &g.vx, maxx)
	update(&g.y, &g.vy, maxy)

	return nil
}

var _ ebiten.Game = &Game{}
