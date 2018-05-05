package basic

import (
	"fmt"
	"testing"
)

type Singer interface {
	Song()
}

type Parrot struct{}

func (p *Parrot) Song() {
	fmt.Println("song ")
}

type Cock struct {
	Singer
}

func (c *Cock) Song() {
	fmt.Println("song with cock")
}

func TestEmbedInterface(t *testing.T) {
	cock := &Cock{Singer: &Parrot{}}
	cock.Song()
	cock = &Cock{}
	cock.Song()
}
