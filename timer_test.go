// Copyright (c) 2014 Maurice Nonnekes <maurice@codeninja.nl>
// All rights reserved.

package timer

import (
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	t0 := New(time.Millisecond * 50)
	t0.Reset()

	time.Sleep(time.Millisecond * 25)

	select {
	case <-t0.C:
		t.Fatal("The timer triggered too early")
	default:
	}

	time.Sleep(time.Millisecond * 25)

	select {
	case <-t0.C:
	default:
		t.Fatal("Expected the timer to have been triggered at this point")
	}
}

func TestResetTimer(t *testing.T) {
	t0 := New(time.Millisecond * 50)
	t0.Reset()

	time.Sleep(time.Millisecond * 25)

	if t0.Reset() == false {
		t.Fatal("Expected to stop the timer while it was still running")
	}

	time.Sleep(time.Millisecond * 25)

	select {
	case <-t0.C:
		t.Fatal("The timer triggered too early")
	default:
	}

	time.Sleep(time.Millisecond * 25)

	select {
	case <-t0.C:
	default:
		t.Fatal("Expected the timer to have been triggered at this point")
	}
}

func TestStopBeforeItExpires(t *testing.T) {
	t0 := New(time.Millisecond * 50)
	t0.Reset()

	time.Sleep(time.Millisecond * 25)

	if t0.Stop() == false {
		t.Fatal("Expected to stop the timer while it was still running")
	}

	time.Sleep(time.Millisecond * 25)

	select {
	case <-t0.C:
		t.Fatal("The timer shouldn't have fired after having called Stop()")
	default:
	}
}

func TestStopAfterItExpired(t *testing.T) {
	t0 := New(time.Millisecond * 50)
	t0.Reset()

	time.Sleep(time.Millisecond * 50)

	if t0.Stop() == true {
		t.Fatal("Expected to stop the timer after it expired")
	}

	select {
	case <-t0.C:
		t.Fatal("The timer shouldn't have fired after having called Stop()")
	default:
	}
}
