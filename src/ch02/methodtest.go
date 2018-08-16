package main

import "fmt"

type myInt int

func (i myInt) add(another int) myInt {
	i = i + myInt(another)
	return i
}

func (i *myInt) add1(another int) myInt {
	*i = *i + myInt(another)
	return *i
}

func (i myInt) add2(another int) int {
	i = i + myInt(another)
	return int(i)
}

func main() {
	i1 := myInt(1)
	i2 := i1.add(2)
	fmt.Println(i1, i2)
	
	i3 := i1.add1(2)
	fmt.Println(i1, i3)
	
	i4 := i1.add2(2)
	fmt.Println(i1, i4)
}