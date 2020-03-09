package defer_test

import "testing"

func TestDefer(t *testing.T) {
	defer t.Log("World")

	t.Log("Hello")
}
