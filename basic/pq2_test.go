package basic

import (
	"container/heap"
	"testing"
	"fmt"
	"container/list"
)

type Item struct {
	Value    string
	Priority int
	index    int
}

type PriorityQueue2 []*Item

func (q PriorityQueue2) Len() int { return len(q) }

func (q PriorityQueue2) Less(i, j int) bool {
	return q[i].Priority < q[j].Priority
}

func (q PriorityQueue2) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *PriorityQueue2) Push(x interface{}) {
	n := len(*q)
	item := x.(*Item)
	item.index = n
	*q = append(*q, item)
}
func (q *PriorityQueue2) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	item.index = -1
	*q = old[0:n-1]
	return item
}

func (q *PriorityQueue2) update(item *Item, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(q, item.index)
}
func (q *PriorityQueue2) Empty() bool {
	return q.Len() == 0
}

func TestPriorityQueue2(t *testing.T) {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}
	q := make(PriorityQueue2, len(items))
	i := 0
	for value, priority := range items {
		q[i] = &Item{
			Value:    value,
			Priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&q)
	item := &Item{
		Value:    "Orange",
		Priority: 1,
	}
	heap.Push(&q, item)
	q.update(item, item.Value, 5)
	for !q.Empty() {
		item := heap.Pop(&q).(*Item)
		fmt.Printf("%.2d:%s(index:%d) ", item.Priority, item.Value, item.index)
	}
}

func TestList(t *testing.T) {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
