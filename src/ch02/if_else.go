package main

import "fmt"

func GetNumber(n int) int {
	if n<=10 {
		return 10
	} else {
		return 11
	}
}

func main() {	
	if a := 5;a<5 {
		fmt.Println("less 5")
	} else {
		fmt.Println("great 5")
	}
	
	b := 6
	fmt.Println("number=", GetNumber(b))
}