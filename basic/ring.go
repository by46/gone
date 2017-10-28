package basic

import (
	"container/ring"
)

func RingLen(r *ring.Ring) int {
	n := 0
	r.Do(func(p interface{}) {
		n += 1
	})
	return n
}
