package asteroids

import (
	"fmt"
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
	wave         int
	messageShown bool
	attractMode  bool
	timer        engine.Timer
}

var _ engine.Interactor = NewWaveMaker()
var _ engine.Ender = NewWaveMaker()

func NewAttractModeWaveMaker() *waveMaker {
	wm := NewWaveMaker()
	wm.attractMode = true
	return wm
}

func NewWaveMaker() *waveMaker {
	return &waveMaker{
		anyAsteroids: false,
		wave:         1,
		timer:        engine.NewTimer(AsteroidDelay),
	}
}

func (w *waveMaker) InteractWith(other any) {
	if _, ok := other.(*asteroid); ok {
		w.anyAsteroids = true
	}
}

func (w *waveMaker) EndUpdate(dt time.Duration, objects *engine.GameObjects) {
	if w.anyAsteroids {
		// there are asteroids present; reset and exit early
		w.anyAsteroids = false
		return
	}

	if !w.attractMode && !w.messageShown {
		objects.Insert(NewMessage(fmt.Sprint("Get ready for Wave ", w.wave), AsteroidDelay-1*time.Second))
		w.messageShown = true
	}

	w.timer.Tick(dt, func() {
		waveSize := InitialWaveSize + (w.wave-1)*WaveSizeIncrement
		if waveSize > MaxWaveSize {
			waveSize = MaxWaveSize
		}

		for i := 0; i < waveSize; i++ {
			objects.Insert(NewAsteroid())
		}
		w.wave++
		w.messageShown = false
	})

	w.anyAsteroids = false
}
