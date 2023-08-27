package asteroids

import (
	"time"

	"github.com/tomasaschan/gosteroids/game/engine"
	"github.com/tomasaschan/gosteroids/game/physics"
)

const (
	SaucerEmergenceTime = 7 * time.Second
	SaucerZigTime       = 1 * time.Second
)

var (
	SaucerInitialVelocity = physics.Point{X: 150, Y: 0}
)

type saucerMaker struct {
	saucerPresent bool
	timer         engine.Timer
}

func NewSaucerMaker() *saucerMaker {
	return &saucerMaker{timer: engine.NewTimer(SaucerEmergenceTime)}
}

func (s *saucerMaker) InteractWith(other any) {
	if s.saucerPresent {
		return
	}

	if _, ok := other.(*saucer); ok {
		s.saucerPresent = true
	}
}

func (s *saucerMaker) Act(dt time.Duration) (result engine.Result) {
	if !s.saucerPresent {
		s.timer.Tick(dt, func() {
			// objects.Insert(NewSaucer(
			// 	physics.Point{X: 0, Y: engine.ScreenSize / 2},
			// 	SaucerInitialVelocity,
			// ))
		})
	}
	s.saucerPresent = false

	return
}

var _ engine.Interactor = NewSaucerMaker()
var _ engine.Actor = NewSaucerMaker()
