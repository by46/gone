package image

import (
	"image"

	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/fcolor"
	"github.com/anthonynsimon/bild/parallel"
)

func Blend(bg image.Image, fg image.Image, fn func(fcolor.RGBAF64, fcolor.RGBAF64) fcolor.RGBAF64) *image.RGBA {
	bgBounds := bg.Bounds()
	fgBounds := fg.Bounds()

	var w, h int
	if bgBounds.Dx() < fgBounds.Dx() {
		w = bgBounds.Dx()
	} else {
		w = fgBounds.Dx()
	}
	if bgBounds.Dy() < fgBounds.Dy() {
		h = bgBounds.Dy()
	} else {
		h = fgBounds.Dy()
	}

	bgSrc := clone.AsRGBA(bg)
	fgSrc := clone.AsRGBA(fg)
	dst := image.NewRGBA(image.Rect(0, 0, w, h))

	parallel.Line(h, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < w; x++ {
				bgPos := y*bgSrc.Stride + x*4
				fgPos := y*fgSrc.Stride + x*4
				result := fn(
					fcolor.NewRGBAF64(bgSrc.Pix[bgPos+0], bgSrc.Pix[bgPos+1], bgSrc.Pix[bgPos+2], bgSrc.Pix[bgPos+3]),
					fcolor.NewRGBAF64(fgSrc.Pix[fgPos+0], fgSrc.Pix[fgPos+1], fgSrc.Pix[fgPos+2], fgSrc.Pix[fgPos+3]))

				result.Clamp()
				dstPos := y*dst.Stride + x*4
				dst.Pix[dstPos+0] = uint8(result.R * 255)
				dst.Pix[dstPos+1] = uint8(result.G * 255)
				dst.Pix[dstPos+2] = uint8(result.B * 255)
				dst.Pix[dstPos+3] = uint8(result.A * 255)
			}
		}
	})

	return dst
}
