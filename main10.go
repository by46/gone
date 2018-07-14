package main

import (
	"fmt"
	"time"
)

var c = make(chan int,1)
var a string

func f() {
	a = "hello world"
	<-c
}
func main() {
	go f()
	time.Sleep(time)
	c <- 0
	fmt.Println(a)
}
