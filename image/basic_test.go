package image

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"testing"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/util"
	"github.com/disintegration/imaging"
	"github.com/stretchr/testify/assert"
)

func TestImageServe(t *testing.T) {
	img, err := imgio.Open("1.jpg")
	assert.Equal(t, nil, err)
	dst := imaging.New(512, 512, color.RGBA{255, 255, 255, 255})
	dst = imaging.PasteCenter(dst, img)

	blue := color.RGBA{0, 0, 255, 255}

	draw.Draw(dst, image.Rect(0, 0, 50, 50), &image.Uniform{blue}, image.ZP, draw.Src)
	imaging.Save(dst, "file.jpg")
}

type Radius struct {
	r    int
	rect image.Rectangle
}

func NewRadius(rect image.Rectangle, r int) *Radius {
	return &Radius{
		r:    r,
		rect: rect,
	}
}

func (r *Radius) ColorModel() color.Model {
	return color.AlphaModel
}

func (r *Radius) Bounds() image.Rectangle {
	return r.rect
}

func (r *Radius) At(x, y int) color.Color {
	xx, yy, rr := x*x, y*y, r.r*r.r
	if xx+yy < rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

func TestImageMask(t *testing.T) {
	dst, err := imgio.Open("1.jpg")
	if err != nil {
		fmt.Printf(err.Error())
	}
	src := image.NewRGBA(dst.Bounds())
	radius := NewRadius(dst.Bounds(), 50)
	draw.DrawMask(src, dst.Bounds(), dst, image.ZP, radius, image.ZP, draw.Over)
	imaging.Save(src, "file2.jpg")
}
func TestGrayImage(t *testing.T) {
	dst, err := imgio.Open("lena.png")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fn := func(c color.RGBA) color.RGBA {
		gray := float64(c.R)*0.3 + float64(c.G)*0.59 + float64(c.B)*0.11
		c.R = uint8(gray)
		c.B = uint8(gray)
		c.G = uint8(gray)
		return c
	}
	dst = adjust.Apply(dst, fn)
	err = imaging.Save(dst, "lena-gray.png")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func TestGrayImageAvg(t *testing.T) {
	dst, err := imgio.Open("lena.png")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fn := func(c color.RGBA) color.RGBA {
		gray := (float64(c.R) + float64(c.G) + float64(c.B)) / 3.0
		c.R = uint8(gray)
		c.B = uint8(gray)
		c.G = uint8(gray)
		return c
	}
	dst = adjust.Apply(dst, fn)
	err = imaging.Save(dst, "lena-gray.png")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func TestGrayImageSaturation(t *testing.T) {
	dst, err := imgio.Open("lena.png")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fn := func(c color.RGBA) color.RGBA {
		h, _, l := util.RGBToHSL(c)
		return util.HSLToRGB(h, 0, l)
	}
	dst = adjust.Apply(dst, fn)
	err = imaging.Save(dst, "lena-gray.png")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func TestGrayImageCustom(t *testing.T) {
	dst, err := imgio.Open("lena.png")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fn := func(c color.RGBA) color.RGBA {
		conversionFactor := 16.0
		value := (float64(c.R) + float64(c.B) + float64(c.G)) / 3.0
		gray := uint8(((value / conversionFactor) + 0.5) * conversionFactor)
		return color.RGBA{R: gray, G: gray, B: gray, A: c.A}
	}
	dst = adjust.Apply(dst, fn)
	err = imaging.Save(dst, "lena-gray.png")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
