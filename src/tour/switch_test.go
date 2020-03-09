package switch_test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSwitch(t *testing.T) {
	t.Log("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		t.Log("OS X.")
	case "linux":
		t.Log("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
