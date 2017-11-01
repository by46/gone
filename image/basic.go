package image

import (
	"github.com/anthonynsimon/bild/imgio"
	"fmt"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/transform"
	"github.com/by46/gone/parameter"
)

func ImageServe() {
	img, err := imgio.Open("1.jpg")
	if err != nil {
		fmt.Printf(err.Error())
	}
	inverted := effect.Invert(img)
	resize := transform.Resize(inverted, 200, 160, transform.Linear)
	rotated := transform.Rotate(resize, 45, nil)
	if err := imgio.Save("filename", rotated, imgio.PNG); err != nil {
		fmt.Printf(err.Error())
	}
	pipe := parameter.NewPipe("w_200,h_100, c_scale")
	if _, err := pipe.Transform(rotated); err != nil {
		fmt.Printf("Error")
	}

}
