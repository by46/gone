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

type Book struct {
	ID int
}

func (c *Cock) Song() {
	fmt.Println("song with cock")
}

func TestEmbedInterface(t *testing.T) {
	cock := &Cock{Singer: &Parrot{}}
	cock.Song()
	cock = &Cock{}
	cock.Song()

	book1 := &Book{ID: 12}

	var book2 *Book
	book2 = new(Book)
	*book2 = *book1
	book2.ID = 15
	fmt.Printf("book1: %v, book2: %v", book1, book2)

	book3 := Book{ID: 30}
	book4 := Book{ID: 40}
	book4 = book3
	book4.ID = 50
	fmt.Printf("book3: %v, book4: %v", book3, book4)
}
