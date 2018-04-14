package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// test map concurrency

var (
	headers     = map[string]string{}
	lock        = &sync.RWMutex{}
	safeHeaders = &sync.Map{}
)

func read() {
	lock.Lock()
	defer lock.Unlock()

	if value, exists := safeHeaders.Load("Name"); exists {
		fmt.Printf("header Name value %s\n", value)
	} else {
		fmt.Printf("Name not exists\n")
	}
}
func write() {
	lock.Lock()
	defer lock.Unlock()
	headers["Name"] = "Benjamin"
	//safeHeaders.Store("Name", "Benjamin")
	fmt.Printf("setting value\n")
}
func main() {
	for i := 0; i < 10000; i ++ {
		seed := rand.Int()
		if seed%2 == 0 {
			go read()
		} else {
			go write()
		}
	}

	fmt.Printf("end loop")

	time.Sleep(time.Second * 60)
}
