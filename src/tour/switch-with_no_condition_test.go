package switch_test

import (
	"testing"
	"time"
)

func TestSwitch(t *testing.T) {
	ti := time.Now()
	switch {
	case ti.Hour() < 12:
		t.Log("Good morning!")
	case ti.Hour() < 17:
		t.Log("Good afternoon.")
	default:
		t.Log("Good evening.")
	}
}
