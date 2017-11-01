package parameter

import (
	"testing"
	"github.com/anthonynsimon/bild/imgio"
	"fmt"
)

func TestPipe(t *testing.T) {
	img, err := imgio.Open("1.jpg")
	if err != nil {
		fmt.Errorf("exception occur, %v", err)
	}
	pipe := NewPipe("w_200,h_150")
	pipe.Transform(img)
}
