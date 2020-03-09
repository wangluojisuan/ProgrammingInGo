package if_test

import (
	"fmt"
	"math"
	"testing"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func TestIt(t *testing.T) {
	t.Log(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
