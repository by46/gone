package effect

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"testing"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/disintegration/imaging"
	"github.com/stretchr/testify/assert"
)

func TestHueEffect(t *testing.T) {
	dst, err := imaging.Open("horses.jpg")
	assert.Nil(t, err)

	dst = adjust.Hue(dst, 40)
	dst = imaging.Resize(dst, 80, 80, imaging.Lanczos)
	err = imgio.Save("horses-hue-40.jpg", dst, imgio.JPEG)
	assert.Nil(t, err)
}

func dumpRect(dst image.Image, message string) {
	fmt.Printf("%s\n", message)
	for i := 0; i < 10; i++ {
		for j := 1000; j < 1010; j++ {
			fmt.Printf("\t%v", dst.At(i, j))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func TestRedEffect(t *testing.T) {
	dst, err := imaging.Open("horses.jpg")
	assert.Nil(t, err)
	//dst = imaging.Resize(dst, 80, 80, imaging.Lanczos)

	fn := func(c color.RGBA) color.RGBA {
		c.G = uint8(math.Min(float64(c.G)*1.50, 255.0))
		return c
	}
	dst = adjust.Apply(dst, fn)
	err = imgio.Save("horses-green2-50.jpg", dst, imgio.JPEG)
	assert.Nil(t, err)
}
func TestRed2Effect(t *testing.T) {
	dst, err := imaging.Open("horses.jpg")
	assert.Nil(t, err)
	dumpRect(dst, "Orignal")

	dst, err = imaging.Open("horses-green-25.jpg")
	assert.Nil(t, err)
	dumpRect(dst, "green:25")

	dst, err = imaging.Open("horses-green-50.jpg")
	assert.Nil(t, err)
	dumpRect(dst, "green:50")

	dst, err = imaging.Open("horses-green-75.jpg")
	assert.Nil(t, err)
	dumpRect(dst, "green:75")

	dst, err = imaging.Open("horses-green-90.jpg")
	assert.Nil(t, err)
	dumpRect(dst, "green:90")
}
