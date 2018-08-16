package main

import "fmt"

func main() {
	var ipSwitches = map[string]bool{}
	
	ipSwitches["192.168.0.1"] = true
	ipSwitches["192.168.1.1"] = false
	
	fmt.Println("ipSwitches=", ipSwitches)
	
	delete(ipSwitches, "192.168.2.1")
	fmt.Println("ipSwitches=", ipSwitches)
	
	delete(ipSwitches, "192.168.0.1")
	fmt.Println("ipSwitches=", ipSwitches)
}