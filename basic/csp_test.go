package basic

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

type Message struct {
	netId int
	data  string
}

type ServeConn struct {
	sendCh   chan Message
	handleCh chan Message
	wg       *sync.WaitGroup
	ctx      context.Context
	cancel   context.CancelFunc
	netId    int
}

func readLoop(conn *ServeConn, wg *sync.WaitGroup) {
	netId, _ := conn.ctx.Value("key").(int)
	ctx, _ := context.WithCancel(conn.ctx)
	cDone := ctx.Done()
	defer wg.Done()
	for {
		time.Sleep(time.Second * 1)
		select {
		case <-cDone:
			fmt.Println("read loop close")
			return
		default:
			conn.handleCh <- Message{netId: netId, data: "hello world"}
		}
	}

}
func handleLoop(conn *ServeConn, wg *sync.WaitGroup) {
	handleCh := conn.handleCh
	sendCh := conn.sendCh
	ctx, _ := context.WithCancel(conn.ctx)

	cDone := ctx.Done()
	defer wg.Done()

	for {
		select {
		case data, ok := <-handleCh:
			if ok {
				data.netId++
				data.data = "I am the whole world"
				sendCh <- data
			}
		case <-cDone:
			fmt.Println("handle loop close")
			return
		}
	}
}
func writeLoop(conn *ServeConn, wg *sync.WaitGroup) {
	sendCh := conn.sendCh
	ctx, _ := context.WithCancel(conn.ctx)
	cDone := ctx.Done()

	defer wg.Done()
	for {
		select {
		case data, ok := <-sendCh:
			if ok {
				fmt.Println(data)
			}

		case <-cDone:
			fmt.Println("write loop close")
			return
		}
	}
}

func TestCSP(t *testing.T) {
	conn := &ServeConn{
		sendCh:   make(chan Message),
		handleCh: make(chan Message),
		wg:       &sync.WaitGroup{},
		netId:    100,
	}
	conn.ctx, conn.cancel = context.WithCancel(context.WithValue(context.Background(), "key", conn.netId))
	loopers := []func(*ServeConn, *sync.WaitGroup){readLoop, handleLoop, writeLoop}
	for _, looper := range loopers {
		conn.wg.Add(1)
		go looper(conn, conn.wg)
	}
	go func() {
		time.Sleep(time.Second * 3)
		conn.cancel()
	}()
	conn.wg.Wait()

}
