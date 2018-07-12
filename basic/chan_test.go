package basic

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	writeCh := make(chan time.Time)
	writeCh <- time.Now()
	close(writeCh)
	<-writeCh
}

func TestChanClose(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)

	stopTimerCh := make(chan struct{})
	time.AfterFunc(time.Second*2, func() { close(stopTimerCh) })
	go func() {
		defer wg.Done()
		select {
		case message, success := <-stopTimerCh:
			fmt.Printf("receive message %v, %v", message, success)
		}
	}()

	wg.Wait()
}
