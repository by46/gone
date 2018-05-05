package basic

import (
	"fmt"
	"testing"
)

type Response struct {
	ID string
}

func (r *Response) Close() {
	fmt.Printf("<Response %s>\n", r.ID)
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func c2() (int) {
	i := 1
	defer func() { i++ }()
	return i
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	fmt.Println("calling g")
	g(0)
	fmt.Println("returned normally from g.")
}
func g(i int) {
	if i > 3 {
		fmt.Println("panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func TestDefer(t *testing.T) {

	responses := []*Response{
		{ID: "1"},
		{ID: "2"},
		{ID: "3"},
		{ID: "3"},
		{ID: "3"},
		{ID: "3"},
		{ID: "3"},
		{ID: "3"},
		{ID: "3"},
		{ID: "3"},
	}

	for _, response := range responses {
		fmt.Printf("visit response %s\n", response.ID)
		defer response.Close()
	}

	fmt.Printf("defer in function %d\n", c())
	fmt.Printf("defer in function %d\n", c2())

	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
	f()
}
