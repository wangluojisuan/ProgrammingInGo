package exercise

import (
	"fmt"
	"testing"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(x - z*z)
	}

	return z
}

func TestIt(t *testing.T) {
	t.Log(Sqrt(3))
}
