package aesf

// create new timer
func NewTimer(delay int64, repeat bool) *Timer {
	timer := new(Timer)
	timer.delay = delay
	timer.repeat = repeat
	return timer
}

type Timer struct {
	delay, acc            int64
	repeat, done, stopped bool
	callback              func()
}

func (t *Timer) SetCallback(c func()) {
	t.callback = c
}

func (t *Timer) SetDelay(delay int64) { t.delay = delay }
func (t *Timer) Delay() int64         { return t.delay }

func (t *Timer) Update(delta int64) {
	if t.done || t.stopped {
		return
	}
	t.acc += delta
	if t.acc >= t.delay {
		t.acc -= t.delay
		if t.repeat {
			t.Reset()
		} else {
			t.done = true
		}
		t.callback()
	}
}

func (t *Timer) Stop() {
	t.stopped = true
}

func (t *Timer) Reset() {
	t.stopped = false
	t.done = false
	t.acc = 0
}

func (t *Timer) Done() bool      { return t.done }
func (t *Timer) IsRunning() bool { return !t.done && !t.stopped && (t.acc < t.delay) }

func (t *Timer) PercentageRemaning() int64 {
	if t.done {
		return 100.
	} else if t.stopped {
		return 0.
	} else {
		return 1. - (t.delay-t.acc)/t.delay
	}
}
