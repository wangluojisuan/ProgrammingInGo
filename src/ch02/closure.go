package main

import (
	"fmt"
)

//func(a, b int, z float64) bool {
//	return a*b < (int)z
//}

func main() {
	var j int = 5
	 a := func(s int) (func (v int)) {
		var i int = s
		return func(v int) {
			fmt.Printf("i, j, v: %d, %d, %d\n", i, j, v)
		}
	 }(20)
	 
	 a(1)
	 
	 j *= 2
	 
	 a(2)
}