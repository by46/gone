package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("outside main")
	go func() {
		fmt.Println("goroutines inside")
	}()

	runtime.Gosched()
	fmt.Println("outside main end")
}
