package basic

import (
	"container/list"
	"sync"

	"github.com/kataras/iris/core/errors"
)

var (
	ErrNoExists = errors.New("key not exists")
)

type CacheNode struct {
	Key, Value interface{}
}

func NewCacheNode(key, value interface{}) *CacheNode {
	return &CacheNode{
		Key:   key,
		Value: value,
	}
}

type CacheLRU struct {
	Capacity int
	ll       *list.List
	m        map[interface{}]*list.Element
	mutex    sync.Mutex
}

func NewCacheLRU(capacity int) *CacheLRU {
	return &CacheLRU{
		Capacity: capacity,
		ll:       list.New(),
		m:        make(map[interface{}]*list.Element),
	}
}

func (lru *CacheLRU) Get(key interface{}) (interface{}, bool) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	if element, ok := lru.m[key]; ok {
		lru.ll.MoveToFront(element)
		return element.Value.(*CacheNode).Value, true
	}
	return nil, false
}
func (lru *CacheLRU) Set(key, value interface{}) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	if element, ok := lru.m[key]; ok {
		lru.ll.MoveToFront(element)
		element.Value.(*CacheNode).Value = value
	}

	element := lru.ll.PushFront(NewCacheNode(key, value))
	lru.m[key] = element

	if lru.ll.Len() > lru.Capacity {
		element := lru.ll.Back()
		if element == nil {
			return
		}
		node := element.Value.(*CacheNode)
		delete(lru.m, node.Key)
		lru.ll.Remove(element)
	}
}

func (lru *CacheLRU) Del(key interface{}) bool {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	if element, ok := lru.m[key]; ok {
		lru.ll.Remove(element)
		delete(lru.m, key)
		return true
	}
	return false
}

func (lru *CacheLRU) Size() int {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	return len(lru.m)
}
