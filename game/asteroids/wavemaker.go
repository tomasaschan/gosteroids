package asteroids

import (
	"time"

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

var _ engine.Beginner = NewWaveMaker()
var _ engine.Interactor = NewWaveMaker()
var _ engine.Ender = NewWaveMaker()

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
	if _, ok := other.(*asteroid); ok {
		w.anyAsteroids = true
	}
}

func (w *waveMaker) EndUpdate(dt time.Duration, objects *engine.GameObjects) {
	if w.anyAsteroids {
		// there are asteroids present; exit early
		return
	}

	w.timer.Tick(dt, func() {
		for i := 0; i < w.waveSize; i++ {
			objects.Insert(NewAsteroid())
		}

		w.waveSize += WaveSizeIncrement
		if w.waveSize > MaxWaveSize {
			w.waveSize = MaxWaveSize
		}
	})
}
