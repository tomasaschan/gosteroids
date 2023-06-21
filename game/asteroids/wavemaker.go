package asteroids

import (
	"time"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/tomasaschan/gosteroids/game/engine"
)

const (
	AsteroidDelay     = 4 * time.Second
	InitialWaveSize   = 2
	WaveSizeIncrement = 2
	MaxWaveSize       = 11
)

type waveMaker struct {
	anyAsteroids bool
	waveSize     int
	timer        engine.Timer
}

var _ engine.Beginner = &waveMaker{}
var _ engine.Interactor = &waveMaker{}
var _ engine.Ender = &waveMaker{}

func NewWaveMaker() *waveMaker {
	return &waveMaker{
		anyAsteroids: false,
		waveSize:     InitialWaveSize,
		timer:        engine.NewTimer(AsteroidDelay),
	}
}

func (w *waveMaker) BeginUpdate() {
	w.anyAsteroids = false
}

func (w *waveMaker) InteractWith(other any) {
	if _, ok := other.(*Asteroid); ok {
		w.anyAsteroids = true
	}
}

func (w *waveMaker) EndUpdate(dt time.Duration, pressedKeys []ebiten.Key, objects *engine.GameObjects) {
	if w.anyAsteroids {
		// there are asteroids present; exit early
		return
	}

	w.timer.Tick(dt, func() {
		for i := 0; i < w.waveSize; i++ {
			a := NewAsteroid()
			objects.Insert(a)
		}

		w.waveSize += WaveSizeIncrement
		if w.waveSize > MaxWaveSize {
			w.waveSize = MaxWaveSize
		}
	})
}
