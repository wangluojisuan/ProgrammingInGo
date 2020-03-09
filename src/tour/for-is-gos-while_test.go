package for_test

import "testing"

func TestFor(t *testing.T) {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	t.Log(sum)
}
