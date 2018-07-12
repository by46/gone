package basic

import (
	"fmt"
	"os"
	"testing"
	"unsafe"
)

type Page struct {
	ID   int
	Size int
}

func (p *Page) String() string {
	return fmt.Sprintf("ID: %d, Size: %d", p.ID, p.Size)
}

func TestPointer(t *testing.T) {
	size := os.Getpagesize()
	buf := make([]byte, 2*size)
	pages := make([]*Page, 0)
	for i := 0; i < 2; i++ {
		p := (*Page)(unsafe.Pointer(&buf[0]))
		p.ID = 12
		p.Size = 64
		pages = append(pages, p)
	}
	for _, p := range pages {
		fmt.Println(p)
	}
}

func TestPointer2(t *testing.T) {
	array := []int{1, 2, 3, 4}
	base := uintptr(unsafe.Pointer(&array[0]))
	size := unsafe.Sizeof(array[0])
	ptr := unsafe.Pointer(base + 2*size)
	element := *(*int)(ptr)
	fmt.Println(element, array[2])
}

func TestPointer3(t *testing.T) {
	array := []int{1}
	fmt.Printf("%v\n", array[:2:2])
}
