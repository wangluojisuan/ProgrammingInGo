package main

import "fmt"

func main() {
	i := 2
	
	switch i {
		case 0:
			fmt.Printf("0")
		case 1:
			fmt.Printf("1")
		case 2:
			fallthrough
		case 3:
			fmt.Printf("3")
		case 4, 5, 6:
			fmt.Printf("4, 5, 6")
		default:
			fmt.Printf("else")
	}
}