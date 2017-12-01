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
	if x >= r.r || y >= r.r {
		return color.White
	}
	x1 := float64(x-r.r) + 0.5
	y1 := float64(y-r.r) + 0.5
	r1 := float64(r.r)
	r2 := float64(r.r + 1)
	xx, yy, rr := x1*x1, y1*y1, r1*r1
	rr2 := r2 * r2
	if xx+yy <= rr {
		return color.White
	} else if xx+yy <= rr2 {

	}
	return color.Transparent
}

func TestImageMask(t *testing.T) {
	dst, err := imgio.Open("lena.png")
	if err != nil {
		fmt.Printf(err.Error())
	}
	src := image.NewRGBA(dst.Bounds())
	radius := NewRadius(dst.Bounds(), 100)
	draw.DrawMask(src, dst.Bounds(), dst, image.ZP, radius, image.ZP, draw.Over)
	imaging.Save(src, "lena-radius2.png")
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

func InCircle(x, y, r int) bool {
	x1 := float64(x) + 0.5
	y1 := float64(y) + 0.5
	r1 := float64(r)
	return x1*x1+y1*y1 < r1*r1
}

func TestPrintCircle(t *testing.T) {
	r := 8
	for i := 1; i <= r; i++ {
		for j := 1; j <= r; j++ {

			fmt.Printf("(%v, %v, %v) ", i, j, InCircle(i, j, r))
		}
		fmt.Println("")
	}
}
