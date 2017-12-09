package image

import (
	"image"
	"image/color"
	"testing"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/blend"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/stretchr/testify/assert"
)

// 正常模式
func TestNormalMix(t *testing.T) {
	dst, err := imgio.Open("lena.png")
	assert.Nil(t, err)
	foreground := image.NewRGBA(dst.Bounds())
	foreground = adjust.Apply(foreground, func(color.RGBA) color.RGBA {
		return color.RGBA{R: 36, G: 1, B: 34, A: 200}
	})
	dst = blend.Normal(dst, foreground)
	err = imgio.Save("lena-normal-mix.png", dst, imgio.PNG)
	assert.Nil(t, err)
}

// 变暗模式
func TestDarkenMix(t *testing.T) {
	dst, err := imgio.Open("lena.png")
	assert.Nil(t, err)
	foreground := image.NewRGBA(dst.Bounds())
	foreground = adjust.Apply(foreground, func(color.RGBA) color.RGBA {
		return color.RGBA{R: 36, G: 1, B: 34, A: 200}
	})
	dst = blend.Darken(dst, foreground)
	err = imgio.Save("lena-darken-mix.png", dst, imgio.PNG)
	assert.Nil(t, err)
}

// 变亮模式
func TestLightenMix(t *testing.T) {
	dst, err := imgio.Open("lena.png")
	assert.Nil(t, err)
	foreground := image.NewRGBA(dst.Bounds())
	foreground = adjust.Apply(foreground, func(color.RGBA) color.RGBA {
		return color.RGBA{R: 36, G: 1, B: 34, A: 200}
	})
	dst = blend.Lighten(dst, foreground)
	err = imgio.Save("lena-lighten-mix.png", dst, imgio.PNG)
	assert.Nil(t, err)
}

// 线性加深
func TestLightBurnMix(t *testing.T) {
	dst, err := imgio.Open("lena.png")
	assert.Nil(t, err)
	foreground := image.NewRGBA(dst.Bounds())
	foreground = adjust.Apply(foreground, func(color.RGBA) color.RGBA {
		return color.RGBA{R: 36, G: 1, B: 34, A: 200}
	})
	dst = blend.LinearBurn(dst, foreground)
	err = imgio.Save("lena-linear-burn-mix.png", dst, imgio.PNG)
	assert.Nil(t, err)
}

// 减去模式
func TestSubtractMix(t *testing.T) {
	dst, err := imgio.Open("lena.png")
	assert.Nil(t, err)
	foreground := image.NewRGBA(dst.Bounds())
	foreground = adjust.Apply(foreground, func(color.RGBA) color.RGBA {
		return color.RGBA{R: 36, G: 1, B: 34, A: 200}
	})
	dst = blend.Subtract(dst, foreground)
	err = imgio.Save("lena-subtract-mix.png", dst, imgio.PNG)
	assert.Nil(t, err)
}

// 正片叠底
func TestMultiplyMix(t *testing.T) {
	dst, err := imgio.Open("lena.png")
	assert.Nil(t, err)
	foreground := image.NewRGBA(dst.Bounds())
	foreground = adjust.Apply(foreground, func(color.RGBA) color.RGBA {
		return color.RGBA{R: 36, G: 1, B: 34, A: 200}
	})
	dst = blend.Multiply(dst, foreground)
	err = imgio.Save("lena-multiply-mix.png", dst, imgio.PNG)
	assert.Nil(t, err)
}

// 滤光模式
func TestScreenMix(t *testing.T) {
	dst, err := imgio.Open("lena.png")
	assert.Nil(t, err)
	foreground := image.NewRGBA(dst.Bounds())
	foreground = adjust.Apply(foreground, func(color.RGBA) color.RGBA {
		return color.RGBA{R: 36, G: 1, B: 34, A: 200}
	})
	dst = blend.Screen(dst, foreground)
	err = imgio.Save("lena-screen-mix.png", dst, imgio.PNG)
	assert.Nil(t, err)
}

// 差值模式
func TestDifferenceMix(t *testing.T) {
	dst, err := imgio.Open("lena.png")
	assert.Nil(t, err)
	foreground := image.NewRGBA(dst.Bounds())
	foreground = adjust.Apply(foreground, func(color.RGBA) color.RGBA {
		return color.RGBA{R: 36, G: 1, B: 34, A: 200}
	})
	dst = blend.Difference(dst, foreground)
	err = imgio.Save("lena-difference-mix.png", dst, imgio.PNG)
	assert.Nil(t, err)
}