package basic

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ch <- true
			count++
			time.Sleep(time.Millisecond)
			count--
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("count: %v", count)
}
