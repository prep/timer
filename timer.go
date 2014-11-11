// Copyright (c) 2014 Maurice Nonnekes <maurice@codeninja.nl>
// All rights reserved.

// Package timer implements an reusable timer in an easy-to-use interface that
// wraps around time.Timer.
package timer

import "time"

type Timer struct {
	C        <-chan time.Time
	timer    *time.Timer
	duration time.Duration
}

// New creates a new timer.
func New(d time.Duration) *Timer {
	timer := time.NewTimer(time.Second)
	timer.Stop()

	return &Timer{C: timer.C, timer: timer, duration: d}
}

// Reset restarts the timer.
// It returns true if the call stops the timer, false if the timer has already
// expired or been stopped.
func (t *Timer) Reset() bool {
	if !t.timer.Reset(t.duration) {
		t.clearChannel()
		return false
	}

	return true
}

// Stop expires the timer.
// It returns true if the call stops the timer, false if the timer has already
// expired or been stopped.
func (t *Timer) Stop() bool {
	if !t.timer.Stop() {
		t.clearChannel()
		return false
	}

	return true
}

// Clear out the timer channel, so it won't trigger
func (t *Timer) clearChannel() {
	select {
	case <-t.timer.C:
	default:
	}
}
