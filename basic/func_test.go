package basic

import (
	"fmt"
	"testing"
)

func greeting(name string, err error) {
	fmt.Printf("greeting hello world %v %v", name, err)
}

func TestFuncWithMultiResult(t *testing.T) {
	greeting(func() (string, error) {
		return "benjamin", fmt.Errorf("greeting abort")
	}())
}
