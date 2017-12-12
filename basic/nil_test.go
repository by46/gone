package basic

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestNilFun(t *testing.T) {
	var dst, src []byte
	dst = nil
	dst = append(dst, src...)
	dst = nil
	copy(dst, src)

	assert.Equal(t, string(dst), "")
	assert.Equal(t, len(dst), 0)
	dst = nil
	dst = dst[:0]
	for i, ch := range dst {
		fmt.Printf("dst %v %v\n", i, ch)
	}

	var student []int
	assert.Equal(t, len(student), 0)
}
