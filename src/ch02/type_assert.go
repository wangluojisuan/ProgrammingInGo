package main

import "fmt"

func main() {
	var v1 int = 1
	var v2, v3 int = 2, 3
	var v4 , v5 = 4, true
	v6, v7 := 6, "hello"
	fmt.Println(v1, v2, v3, v4, v5, v6, v7)
}