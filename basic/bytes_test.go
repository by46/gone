package basic

import (
	"fmt"
	"testing"
)

func process(buf []byte){
	buf[0] = 0x30
}

func TestBytesEqual(t *testing.T) {
	names := []byte("hello world")
	process(names)
	fmt.Printf("names: %s\n", names[:])
}
