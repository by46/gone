package basic

import (
	"testing"
	"fmt"
)

type demo1 struct {
	int
}
type demo2 struct {
	float32
	*demo1
}

func TestEmbedStruct(t *testing.T) {
	d := &demo2{
		float32: 12.0,
		demo1:   &demo1{1,},
	}

	fmt.Printf("demo %v\n", d.demo1.int)
	fmt.Printf("demo2 %v\n", d.int)
}
