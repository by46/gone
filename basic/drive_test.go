package basic

import (
	"testing"
	"fmt"
)

type Pipe map[string]string

func (p *Pipe) Add() {
	(*p)["name"] = "123"
}
func NewPipe() *Pipe {
	return &Pipe{}
}
func (p *Pipe) Foreach(action func(key, value string)) {

	for key, value := range *p {
		action(key, value)
	}
}
func TestDrive(t *testing.T) {
	x := NewPipe()
	x.Add()
	fmt.Printf("%v\n", x)
	fn := func(key, value string) {
		fmt.Printf("%s:%s\n", key, value)
	}
	x.Foreach(fn)
}
