package basic

import (
	"testing"
	"container/heap"
	"math/rand"
	"fmt"
	"time"
)

type PriorityQueue []int

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i] < (*pq)[j]
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Push(v interface{}) {
	*pq = append(*pq, v.(int))
}

func (pq *PriorityQueue) Pop() (v interface{}) {
	*pq, v = (*pq)[:pq.Len()-1], (*pq)[pq.Len()-1]
	return
}

func (pq *PriorityQueue) Empty() bool {
	return pq.Len() == 0
}

func (pq *PriorityQueue) Deque() (v interface{}) {
	return heap.Pop(pq)
}
func (pq *PriorityQueue) Enqueue(v interface{}) {
	heap.Push(pq, v)
}

func TestInit1(t *testing.T) {
	queue := new(PriorityQueue)
	rand.Seed(time.Now().Unix())
	//heap.Init(queue)
	for i := 1; i < 20; i++ {
		queue.Enqueue(rand.Intn(100))
	}
	for !queue.Empty() {
		fmt.Printf("First Priority Element: %v\n", queue.Deque())
	}
}
