package engine

import "time"

type TimerAction func()

type Timer struct {
	period  time.Duration
	elapsed time.Duration
}

func NewTimer(period time.Duration) Timer {
	return Timer{period: period}
}

func (t *Timer) Tick(dt time.Duration, actionIfElapsed TimerAction) {
	if t.elapsed+dt < t.period {
		t.elapsed += dt
		return
	}

	actionIfElapsed()
	t.elapsed = 0 * time.Second
}
