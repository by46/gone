package main

import "time"

func Read() {
	writeCh := make(chan time.Time)
	writeCh <- time.Now()
	//close(writeCh)
	<-writeCh
}

func main() {
	go Read()

	time.Sleep(2 * time.Second)
}
