package basic

import (
	"container/list"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type Conn struct {
	Name string
}

type KeyLRU struct {
	idleMux sync.Mutex
	ll      *list.List
	m       map[*Conn]*list.Element
}

func (k *KeyLRU) add(c *Conn) {
	k.idleMux.Lock()
	defer k.idleMux.Unlock()

	if k.ll == nil {
		k.ll = list.New()
		k.m = make(map[*Conn]*list.Element)
	}

	ele := k.ll.PushFront(c)
	if _, ok := k.m[c]; ok {
		panic("the connection already in list")
	}
	k.m[c] = ele
}

func (k *KeyLRU) remove(c *Conn) {
	k.idleMux.Lock()
	defer k.idleMux.Unlock()
	if ele, ok := k.m[c]; ok {
		k.ll.Remove(ele)
		delete(k.m, c)
	}
}

func (k *KeyLRU) removeOldest() *Conn {
	k.idleMux.Lock()
	defer k.idleMux.Unlock()
	ele := k.ll.Back()
	c := ele.Value.(*Conn)
	delete(k.m, c)
	return c
}

func (k *KeyLRU) len() int {
	k.idleMux.Lock()
	defer k.idleMux.Unlock()
	return len(k.m)
}

func TestLRU(t *testing.T) {
	var wg sync.WaitGroup
	lru := &KeyLRU{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(no int) {
			defer wg.Done()
			c := &Conn{Name: fmt.Sprintf("worker %d", no)}
			lru.add(c)
			n := rand.Intn(100)
			fmt.Printf("worker %d sleep %d\n", no, n)
			time.Sleep(time.Duration(n) * time.Millisecond)
			lru.remove(c)
		}(i)
	}
	wg.Wait()
}

func TestCacheLRU(t *testing.T) {
	cache := NewCacheLRU(30)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			for j := 0; j < 10; j++ {
				key := rand.Intn(1000)
				value := rand.Intn(1000)
				cache.Set(key, value)
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("LRU size: %d", cache.Size())
}

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

func TestHappenBefore(t *testing.T) {
	var c = make(chan int, 10)
	var a string
	f := func() {
		time.Sleep(time.Millisecond*100)
		a = "hello world"
		c <- 0
	}
	go f()
	<-c
	fmt.Println(a)
	time.Sleep(time.Second)
}
