package image

import (
	"image/color"
	"os"
	"testing"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/disintegration/imaging"
	"github.com/stretchr/testify/assert"
)

func TestTransparent2(t *testing.T) {
	opacity := 0.0
	fn := func(c color.RGBA) color.RGBA {
		if c.A == 0 {
			c = color.RGBA{R: 255, G: 255, B: 255, A: 255}
		}
		return color.RGBA{
			R: uint8(float64(c.R) * opacity),
			G: uint8(float64(c.G) * opacity),
			B: uint8(float64(c.B) * opacity),
			A: uint8(255.0 * opacity),
		}
	}
	dst, err := imaging.Open("nasa.jpg")
	assert.Nil(t, err)
	writer, err := os.OpenFile("lena-transparent2.jpeg", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer writer.Close()
	assert.Nil(t, err)
	dst = adjust.Apply(dst, fn)
	err = imaging.Encode(writer, dst, imgio.PNG)
	assert.Nil(t, err)
}
