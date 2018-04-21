package basic

import (
	"testing"
	"sync"
	"fmt"
)

func TestSyncOnce(t *testing.T) {
	once := new(sync.Once)

	init := func() {
		fmt.Printf("init method\n")
	}
	for i := 0; i <= 10; i++ {
		fmt.Print(i)
		once.Do(init)
	}
}
