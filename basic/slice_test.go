package basic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceType(t *testing.T) {

	person := make([]int, 0)
	val := reflect.ValueOf(person)
	fmt.Printf("type %v", val.Type())
}

func TestSliceType2(t *testing.T) {
	person := make([]int, 0)
	for _, v := range person {
		fmt.Printf("hello %v", v)
	}

	var person2 []int
	fmt.Printf("person2 %v", person2)
	for _, v := range person2 {
		fmt.Printf("element in person2 %v", v)
	}
}
