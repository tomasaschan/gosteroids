package asteroids

import (
	"time"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/tomasaschan/gosteroids/game/engine"
	"github.com/tomasaschan/gosteroids/game/physics"
)

const (
	SaucerEmergenceTime = 7 * time.Second
	SaucerZigTime       = 1 * time.Second
)

var (
	SaucerInitialVelocity = physics.Point{150, 0}
)

type saucerMaker struct {
	saucerPresent bool
	timer         engine.Timer
}

func NewSaucerMaker() *saucerMaker {
	return &saucerMaker{timer: engine.NewTimer(SaucerEmergenceTime)}
}

func (s *saucerMaker) BeginUpdate() {
	s.saucerPresent = false
}

func (s *saucerMaker) InteractWith(other any) {
	if s.saucerPresent {
		return
	}

	if _, ok := other.(*saucer); ok {
		s.saucerPresent = true
	}
}

func (s *saucerMaker) EndUpdate(dt time.Duration, pressedKeys []ebiten.Key, objects *engine.GameObjects) {
	if !s.saucerPresent {
		s.timer.Tick(dt, func() {
			objects.Insert(NewSaucer(
				physics.Point{X: 0, Y: engine.ScreenSize / 2},
				SaucerInitialVelocity,
			))
		})
	}
}

var _ engine.Beginner = &saucerMaker{}
var _ engine.Interactor = &saucerMaker{}
var _ engine.Ender = &saucerMaker{}
