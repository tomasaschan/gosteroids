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
	safetyMargin   = 20
	missileTimeout = 2 * time.Second
	missileSpeed   = 400
)

type missile struct {
	state     physics.State
	key       string
	age       time.Duration
	colliding bool
}

func NewMissile(origin physics.State, shooterRadius float64, key string) *missile {
	state := origin
	state.P = state.P.Add(physics.P(shooterRadius+missileRadius+safetyMargin, 0).Rotate(state.Theta))
	state.V = state.V.Add(physics.P(missileSpeed, 0).Rotate(state.Theta))
	state.Vtheta = 0
	return &missile{state: state, key: key}
}

var _ engine.Interactor = NewMissile(physics.State{}, 0, "")
var _ engine.Ender = NewMissile(physics.State{}, 0, "")
var _ engine.Drawable = NewMissile(physics.State{}, 0, "")

func (m *missile) EndUpdate(dt time.Duration, objects *engine.GameObjects) {
	if m.age > missileTimeout || m.colliding {
		objects.Remove(m)
	}

	m.colliding = false
	m.state.Evolve(dt)
	m.age += dt
}

func (m *missile) InteractWith(other any) {
	if asteroid, ok := other.(*asteroid); ok {
		m.colliding = m.colliding || AreColliding(m.state, asteroid.State, missileRadius+asteroid.radius)
	}
	if ship, ok := other.(*Ship); ok {
		m.colliding = m.colliding || AreColliding(m.state, ship.State, shipRadius+missileRadius)
	}
}

func (m *missile) Draw(target pixel.Target) {
	imd := imdraw.New(nil)

	imd.Push(pixel.Vec(m.state.P))
	imd.Circle(4, 0)

	imd.Draw(target)
}
