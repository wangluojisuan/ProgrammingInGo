package if_test

import (
	"fmt"
	"math"
	"testing"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func TestIf(t *testing.T) {
	t.Log(sqrt(2))
	t.Log(sqrt(-4))
}
