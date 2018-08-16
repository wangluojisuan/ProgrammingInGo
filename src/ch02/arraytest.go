package main

import "fmt"

func main() {
	var ipv4 [4]int = [4]int{192, 168, 0, 1}	
	fmt.Println(ipv4)
	
	ipv4_1 := [4]int{192, 168, 1, 1}
	fmt.Println(ipv4_1)
	
	var ipv4_2 [4]int = [...]int{192, 168, 2, 1}
	fmt.Println(ipv4_2)
	
	ipv4_3 := [...]int{192, 168, 3, 1}
	fmt.Println(ipv4_3, len(ipv4_3), cap(ipv4_3))	
}