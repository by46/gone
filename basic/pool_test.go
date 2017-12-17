package basic

import (
	"testing"
	"sync"
	"fmt"
	"runtime"
)

type PoolItem struct {
	Name string
}

func TestSyncPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	a := pool.Get()
	pool.Put(1)
	b := pool.Get()
	fmt.Print(a, b)
}

func TestSyncPool2(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	a := pool.Get()
	pool.Put(1)
	runtime.GC()
	b := pool.Get()
	fmt.Print(a, b)
}
func TestSyncPool3(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return &PoolItem{}
		},
	}
	s1 := pool.Get().(*PoolItem)
	s1.Name = "benjamin"
	runtime.GC()
	fmt.Print(s1)
}
