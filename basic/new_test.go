package basic

import (
	"fmt"
	"testing"
)

type Student struct {
	Name string
}

const (
	Enone  = 1
	Eio    = 3
	Einval = 4
)

func TestNewFunc(t *testing.T) {
	s := new(Student)
	fmt.Printf("Name: %s\n", s.Name)
	s2 := []string{
		1: "no error", Eio: "Eio", Einval: "invalid argument"}
	fmt.Printf("Name: %v %v %v\n", s2, len(s2), cap(s2))
	s2 = make([]string, 0, 5)
	s2 = append(s2, "benjamin")
	s2 = append(s2, "benjamin")
	s2 = append(s2, "benjamin")
	s2 = append(s2, "benjamin")
	s2 = append(s2, "benjamin")
	s2 = append(s2, "benjamin")
	fmt.Printf("Name: %v %v %v\n", s2, len(s2), cap(s2))

	var s3 []string
	s3 = s2
	s2 = append(s2, "wendy")
	fmt.Printf("Name: %v %v %v\n", s2, len(s2), cap(s2))
	fmt.Printf("Name: %v %v %v\n", s3, len(s3), cap(s3))

	s4 := s2[0:10]
	s2[0] = "Error"
	fmt.Printf("Name: %v %v %v\n", s4, len(s4), cap(s4))

}

func TestSlice(t *testing.T) {
	s1 := make([]int, 5, 10)
	s2 := s1[0:5]
	s1 = append(s1, 6)
	fmt.Printf("s1 %v\n", s1)
	s2 = append(s2, 7)
	fmt.Printf("s1 %v\n", s1)
	fmt.Printf("s2 %v\n", s2)
}

func TestOverflow(t *testing.T) {
	var r, g, b uint8
	r = 255
	g = 255
	b = 255
	fmt.Printf("(r+g+b) %v \n", float64(r+g+b))
}
