package main

import "fmt"

type PersonInfo struct{
	ID string
	Name string
	Address string
}

func main(){
	var personInfoDB map[string] PersonInfo
	personInfoDB = make(map[string] PersonInfo)
	
	// 插入数据
	personInfoDB["12345"] = PersonInfo{"12345", "Tom", "Room 203"}
	personInfoDB["1"] = PersonInfo{"1", "Jack", "Room 101"}
	
	// 查找数据
	person, ok := personInfoDB["12345"]
	
	if ok{
		fmt.Println("Found person", person.Name, "with ID 12345.")
	} else {
		fmt.Println("Did not find person with 12345.")
	}
}