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
	var once sync.Once
	doCancel := func() {
		once.Do(func() {
			close(stopTimerCh)
		})
	}
	timer := time.NewTimer(time.Second * 2)
	time.AfterFunc(time.Second*2, doCancel)
	go func() {
		defer wg.Done()
		select {
		case <-timer.C:
			fmt.Printf("do cancel because time")
			doCancel()
		case message, success := <-stopTimerCh:
			fmt.Printf("receive message %v, %v", message, success)
			timer.Stop()
		}
	}()

	wg.Wait()
}

func TestChanNil(t *testing.T) {
	var ch chan int
	ch = nil
	time.AfterFunc(time.Second, func() {
		if ch != nil {
			close(ch)
		}
	})
	select {
	case <-time.NewTimer(time.Millisecond * 100).C:
		fmt.Printf("timeout\n")
	case message, success := <-ch:
		fmt.Printf("receive %v %v", message, success)
	}
	fmt.Printf("success")
}
