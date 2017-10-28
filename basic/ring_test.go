package basic

import (
	"testing"
	"container/ring"
	"fmt"
)

func TestRingLen(t *testing.T) {
	r := ring.New(5)
	n := RingLen(r)
	if n != 5 {
		t.Errorf("r.Len() == %d, expected %d", n, 5)
	}
}

func TestInsertElements(t *testing.T) {
	r := ring.New(1)
	r.Value = 1

	r.Link(&ring.Ring{Value: 2})
	r.Link(&ring.Ring{Value: 5})
	r.Link(&ring.Ring{Value: 7})
	(&ring.Ring{Value: 10}).Link(r)
	r.Do(func(p interface{}) {
		fmt.Printf("Current element %d\n", p.(int))
	})
}

func TestSuicide(t *testing.T) {
	count := 41
	r := ring.New(count)
	for i := 1; i <= count; i++ {
		r.Value = i
		r = r.Next()
	}
	r.Unlink(2)
	iterateRing(r)

	for ; r.Len() > 2; {
		r = r.Move(1)
		r.Unlink(1)
		r = r.Move(1)
	}
	iterateRing(r)
}

func iterateRing(r *ring.Ring) {
	fmt.Printf("Elements sequence: ")
	r.Do(func(p interface{}) {
		if p == nil {
			fmt.Printf("%v ", p)
		} else {
			fmt.Printf("%d ", p.(int))
		}
	})
	fmt.Printf("\n")
}
func makeN(r *ring.Ring, base int) {
	if r == nil {
		return
	}
	p := r
	for i := 1; i == 1 || p != r; r = r.Next() {
		r.Value = base + i
		i++
	}
}

func TestJoinTwoRing(t *testing.T) {
	l := ring.New(5)
	makeN(l, 0)
	r := ring.New(10)
	makeN(r, 100)
	l.Link(r)
	iterateRing(l)
	iterateRing(r)
	l2 := ring.New(10)
	makeN(l2, 100)
	l2Removed := l2.Unlink(3)
	iterateRing(l2)
	iterateRing(l2Removed)
}
