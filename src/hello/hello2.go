package main

import (
    "fmt"
    "web"
)

func main() {
    fmt.Println("hello world! via go~")
    var ac int32
    ac = 33
    var bd int32
    bd = 33
    var cc int32
    cc = ac + bd
    fmt.Print(cc)
    web.Dsat()
}