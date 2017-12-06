package image

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/anthonynsimon/bild/util"
	"github.com/disintegration/imaging"
)

const (
	RadiusWidth = 0.8
)

var (
	Mapping  = make(map[int]color.RGBA)
	Mapping2 = make(map[int]int)
)

// toNRGBA converts any image type to *image.NRGBA with min-point at (0, 0).
func toNRGBA(img image.Image) *image.NRGBA {
	if img, ok := img.(*image.NRGBA); ok && img.Bounds().Min.Eq(image.ZP) {
		return img
	}
	return imaging.Clone(img)
}

func InCorner(x, y, r int) float64 {

	originX, originY := r, r
	x1 := float64(x-originX) + 0.5
	y1 := float64(y-originY) + 0.5
	r1 := float64(r)
	r2 := float64(r) + RadiusWidth
	xx, yy, rr := x1*x1, y1*y1, r1*r1
	rr2 := r2*r2
	if xx+yy < rr {
		Mapping2[x] = y
		return 1.0
	} else if xx+yy < rr2 {
		//y1 := Mapping2[x]
		//return math.Max(1.0-float64(y1-y)*0.2, 0.2)
		offset := math.Max((r2-math.Sqrt(xx+yy))/RadiusWidth, 0.0)
		return float64(uint8(offset*10.0)/2) * 0.1
		if offset < 0.1 {
			return 0.1
		} else if offset < 0.2 {
			return 0.6
		} else if offset < 0.6 {
			return 0.7
		} else if offset < 0.8 {
			return 0.8
		} else {
			return 0.9
		}
	}
	return 0.0
}

func Radius1(img image.Image, radius int) image.Image {
	dst := toNRGBA(img)
	for i := radius; i >= 0; i-- {
		for j := radius; j >= 0; j-- {
			opacity := InCorner(i, j, radius)
			offset := dst.PixOffset(i, j)
			r := dst.Pix[offset]
			g := dst.Pix[offset+1]
			b := dst.Pix[offset+2]
			a := dst.Pix[offset+3]
			c := color.RGBA{r, g, b, a}
			if opacity == 1.0 {
				Mapping[i] = c
			} else {
				if opacity == 0.0 {
					dst.Pix[offset] = 0
					dst.Pix[offset+1] = 0
					dst.Pix[offset+2] = 0
					dst.Pix[offset+3] = 0
					continue
				}
				c = Mapping[i]
				h, s, l := util.RGBToHSL(c)
				l += 0.15
				c = util.HSLToRGB(h, s, math.Min(l, 0.9))
				fmt.Printf("(%v, %v) %v\n", i, j, opacity)
				dst.Pix[offset] = uint8(float64(c.R) * opacity)
				dst.Pix[offset+1] = uint8(float64(c.G) * opacity)
				dst.Pix[offset+2] = uint8(float64(c.B) * opacity)
				dst.Pix[offset+3] = uint8(255.0 * opacity)
			}
		}
	}
	return dst
}
