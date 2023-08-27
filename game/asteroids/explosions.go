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

type LineFragment struct {
	State physics.State
	a, b  pixel.Vec
	age   time.Duration
}

var _ engine.Drawable = &LineFragment{}
var _ engine.Actor = &LineFragment{}

func (f *LineFragment) Draw(target pixel.Target) {
	imd := imdraw.New(nil)
	imd.SetMatrix(pixel.IM.Rotated(pixel.ZV, f.State.Theta).Moved(pixel.Vec(f.State.P)))
	imd.Push(f.a, f.b)
	imd.Line(4)

	imd.Draw(target)
}

func (f *LineFragment) Act(dt time.Duration) engine.Result {
	f.age += dt

	if f.age > FragmentLifetime {
		return engine.Result{RemoveSelf: true}
	}

	f.State.Evolve(dt)
	return engine.Result{}
}

func NewLineFragment(state physics.State, a, b pixel.Vec) *LineFragment {
	state.Vtheta += 2*rand.Float64() - 1
	state.V = state.V.Add(randomAdjustment())
	return &LineFragment{
		State: state,
		a:     a,
		b:     b,
	}
}

func randomAdjustment() physics.Point {
	magnitude := rand.Float64() * 25
	angle := rand.Float64() * 2 * math.Pi

	return physics.P(magnitude, 0).Rotate(angle)
}
