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
