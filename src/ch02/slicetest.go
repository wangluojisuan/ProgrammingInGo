package main

import "fmt"

func main() {
	var ips = []string{"192.168.0.1", "192.168.0.2", "192.168.0.3"}
	
	fmt.Println(ips, "len=", len(ips), "cap=", cap(ips))
	
	ips1 := append(ips, "192.168.0.4")
	fmt.Println("ips=", ips)
	fmt.Println("ips1=", ips1)
	
	fmt.Println("****************************")
	ips1[0] = "0.0.0.0"
	fmt.Println("ips=", ips)
	fmt.Println("ips1=", ips1)
	
	var ips2 = make([]string, 2)
	var ips3 = make([]string, 2, 4)
	ips2[0] = "192.168.2.1"
	ips2[1] = "192.168.2.2"
	fmt.Println("ips2=", ips2, "len=", len(ips2), "cpa=", cap(ips2))
	ips3[0] = "192.168.3.1"
	ips3[1] = "192.168.3.2"
	fmt.Println("ips3=", ips3, "len=", len(ips3), "cap=", cap(ips3))
	
	ips3_1 := append(ips3, "192.168.3.31");
	ips3[0] = "0.0.0.0"
	fmt.Println("ips3=", ips3, "len=", len(ips3), "cap=", cap(ips3))
	fmt.Println("ips3_1=", ips3_1, "len=", len(ips3_1), "cap=", cap(ips3_1))
}