package defer_test

import "testing"

func TestDefer(t *testing.T) {
	t.Log("counting")

	for i := 0; i < 10; i++ {
		defer t.Log(i)
	}

	t.Log("done")
}
