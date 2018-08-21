package main

import "fmt"
import "errors"

type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

type function func(string) string

type myTalk string
type otherTalk int

func (talk *myTalk)Hello(userName string) string {
	tmp := "Hello " + userName
	*talk = myTalk(tmp)	
	return tmp
}

func (talk *myTalk)Talk(heard string) (saying string, end bool, err error) {
	saying = ""
	end = false
	err = errors.New("empty")
	return
}

func (talk otherTalk) Hello(userName string) string {
	return "Hello " + userName
}

func (talk otherTalk)Talk(heard string) (saying string, end bool, err error) {
	saying = ""
	end = false
	err = errors.New("empty")
	return
}

func (t Talk, f function) string {
	if f == nil {
		return ""
	}
	
	return t.Hello("")
}


func main() {
	myTalkValue := myTalk("")
	//otherTalkValue := myTalkValue
	
	hi := myTalkValue.Hello("World")
	fmt.Println("Value=", myTalkValue, ";Ret=", hi)
}