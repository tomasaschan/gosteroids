package asteroids

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/tomasaschan/gosteroids/game/engine"
	"github.com/tomasaschan/gosteroids/game/physics"
)

const (
	missileRadius  = 5
	missileTimeout = 4 * time.Second
	missileSpeed   = 500 / 3
)

type missile struct {
	state physics.State
	age   time.Duration
}

func NewMissile(origin physics.State, shooterRadius float64) *missile {
	state := origin
	state.P = state.P.Add(physics.P(shooterRadius+missileRadius+1, 0).Rotate(state.Theta))
	state.V = state.V.Add(physics.P(missileSpeed, 0).Rotate(state.Theta))
	state.Vtheta = 0
	return &missile{state: state}
}

var _ engine.Interactor = NewMissile(physics.State{}, 0)
var _ engine.Ender = NewMissile(physics.State{}, 0)
var _ engine.Drawable = NewMissile(physics.State{}, 0)

func (m *missile) EndUpdate(dt time.Duration, objects *engine.GameObjects) {
	if m.age > missileTimeout {
		objects.Remove(m)
	}

	m.state.Evolve(dt)
	m.age += dt
}

func (*missile) InteractWith(any) {

}

func (m *missile) Draw(target pixel.Target) {
	imd := imdraw.New(nil)

	imd.Push(pixel.Vec(m.state.P))
	imd.Circle(4, 0)

	imd.Draw(target)
}
