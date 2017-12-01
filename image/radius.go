package image

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/disintegration/imaging"
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
	r2 := float64(r + 1)
	xx, yy, rr := x1*x1, y1*y1, r1*r1
	rr2 := r2 * r2
	if xx+yy <= rr {
		return 1.0
	} else if xx+yy <= rr2 {
		offset := (r2 - math.Sqrt(xx+yy)) / 1
		if offset < 0.2 {
			return 0.5
		} else if offset < 0.4 {
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
	for i := 0; i < radius; i++ {
		for j := 0; j < radius; j++ {
			opacity := InCorner(i, j, radius)
			if opacity < 1.0 {
				fmt.Printf("(%v, %v) %v\n", i, j, opacity)
				c := dst.At(i, j)
				c1, _ := c.(color.NRGBA)
				c1 = color.NRGBA{
					R: uint8(float64(c1.R) * opacity),
					G: uint8(float64(c1.G) * opacity),
					B: uint8(float64(c1.B) * opacity),
					A: uint8(255.0 * opacity),
				}
				dst.Set(i, j, c1)
			}
		}
	}
	return dst
}
