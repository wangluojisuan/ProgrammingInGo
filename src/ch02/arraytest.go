package main

import "fmt"

func main() {
	var ipv4 [4]int = [4]int{192, 168, 1, 1}	
	fmt.Println(ipv4)
	
	ipv4_1 := [4]int{192, 168, 2, 1}
	fmt.Println(ipv4_1)
}