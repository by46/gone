package image

import (
	"image"
	"image/color"
	"os"
	"testing"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/imgio"
	_ "github.com/anthonynsimon/bild/util"
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

func TestFill(t *testing.T) {
	dst, err := imaging.Open("nasa.jpg")
	assert.Nil(t, err)
	writer, err := os.OpenFile("lena-fill.jpeg", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer writer.Close()
	assert.Nil(t, err)
	dst = imaging.Fill(dst, 1200, 900, imaging.Center, imaging.Lanczos)
	err = imaging.Encode(writer, dst, imgio.JPEG)
	assert.Nil(t, err)
}

func TestFit(t *testing.T) {
	dst, err := imaging.Open("nasa.jpg")
	assert.Nil(t, err)
	writer, err := os.OpenFile("lena-fit.jpeg", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer writer.Close()
	assert.Nil(t, err)
	dst = imaging.Fit(dst, 250, 200, imaging.Lanczos)
	err = imaging.Encode(writer, dst, imgio.JPEG)
	assert.Nil(t, err)
}

func TestResize(t *testing.T) {
	dst, err := imaging.Open("lena.png")
	assert.Nil(t, err)
	writer, err := os.OpenFile("lena-resize2.jpeg", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer writer.Close()
	assert.Nil(t, err)
	dst = imaging.Fit(dst, 300, 200, imaging.Lanczos)
	err = imaging.Encode(writer, dst, imgio.JPEG)
	assert.Nil(t, err)
}

func setColor(input *image.RGBA, col color.RGBA) *image.RGBA {
	fn := func(c color.RGBA) color.RGBA {
		return col
	}
	return adjust.Apply(input, fn)
}
func setOpacity(c color.RGBA, opacity float64) color.RGBA {
	//h, s, l := util.RGBToHSL(c)
	//s = s * opacity
	//return util.HSLToRGB(h, s, l)
	c.R = uint8(float64(c.R) * opacity)
	c.G = uint8(float64(c.G) * opacity)
	c.B = uint8(float64(c.B) * opacity)
	c.A = uint8(255.0 * opacity)
	return c
}

func TestTransparent3(t *testing.T) {
	dst := image.NewNRGBA(image.Rect(0, 0, 1000, 200))
	cube := image.NewRGBA(image.Rect(0, 0, 100, 100))
	c := color.RGBA{48, 72, 20, 255}
	for i := 1; i <= 10; i += 1 {
		r := float64(i) * 0.1
		cube = setColor(cube, setOpacity(c, r))
		dst = imaging.Paste(dst, cube, image.Pt(100*(i-1), 0))
	}
	err := imaging.Save(dst, "transparent.png")
	assert.Nil(t, err)
}
