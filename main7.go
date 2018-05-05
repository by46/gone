package main

import (
	"fmt"
	"time"
)

type Options struct{}

func Read() {
	writeCh := make(chan *Options, )
	writeCh <- new(Options)
	close(writeCh)
	for {
		select {
		case x := <-writeCh:

			fmt.Println("result ", x)
		default:
			fmt.Println("default")
		}
	}
}

func main() {
	go Read()

	time.Sleep(20 * time.Second)
}
