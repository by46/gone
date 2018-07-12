package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(no int) {
			ch <- true
			count++
			fmt.Printf("worker %v cont %v\n", no, count)
			time.Sleep(time.Millisecond)
			fmt.Printf("end worker %v cont %v\n", no, count)
			count--
			<-ch
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("count: %v", count)
}
