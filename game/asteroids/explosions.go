package asteroids

import (
	"math"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/tomasaschan/gosteroids/game/engine"
	"github.com/tomasaschan/gosteroids/game/physics"
)

const FragmentLifetime = 5 * time.Second

type lineFragment struct {
	state  physics.State
	length float64
	age    time.Duration
}

var _ engine.Drawable = &lineFragment{}
var _ engine.Ender = &lineFragment{}

func (f *lineFragment) Draw(target pixel.Target) {
	imd := imdraw.New(nil)
	imd.SetMatrix(pixel.IM.Rotated(pixel.ZV, f.state.Theta).Moved(pixel.Vec(f.state.P)))
	imd.Push(pixel.V(-f.length/2, 0), pixel.V(f.length/2, 0))
	imd.Line(4)

	imd.Draw(target)
}

func (f *lineFragment) EndUpdate(dt time.Duration, objects *engine.GameObjects) {
	f.age += dt

	if f.age > FragmentLifetime {
		objects.Remove(f)
		return
	}

	f.state.Evolve(dt)
}

func LineFragment(state physics.State, a, b physics.Point) *lineFragment {
	relativeMidpoint := b.Add(a).Scale(.5)
	location := state.P.Add(relativeMidpoint)

	velocity := state.V.Add(randomAdjustment())

	dx, dy := b.X-a.X, b.Y-a.Y
	theta := state.Theta + math.Atan2(dy, dx)

	return &lineFragment{
		state: physics.State{
			P:      location,
			V:      velocity,
			Theta:  theta,
			Vtheta: math.Pi / 4,
		},
		length: b.Sub(a).Length(),
	}
}

func randomAdjustment() physics.Point {
	magnitude := rand.Float64() * 25
	angle := rand.Float64() * 2 * math.Pi

	return physics.P(magnitude, 0).Rotate(angle)
}
