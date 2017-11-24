package image

import (
	"image"
	"testing"

	"image/color"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/disintegration/imaging"
	"github.com/pressly/goico"
	"github.com/stretchr/testify/assert"
)

func init() {

	image.RegisterFormat("ico", "\x00\x00\x01\x00?????\x00", ico.Decode, ico.DecodeConfig)
}

func TestTransparent(t *testing.T) {
	img, err := imgio.Open("favicon.ico")
	assert.Nil(t, err)

	dst := imaging.New(64, 64, color.RGBA{0, 0, 0, 0,})

	dst = imaging.PasteCenter(dst, img)
	err = imgio.Save("favicon", dst, imgio.PNG)
	assert.Nil(t, err)
}

func TestTransparentLena(t *testing.T) {
	img, err := imgio.Open("lena.png")
	assert.Nil(t, err)

	dst := imaging.New(800, 600, color.RGBA{0, 0, 0, 0,})
	dst = imaging.PasteCenter(dst, img)
	err = imgio.Save("lena-transparent", dst, imgio.PNG)
	assert.Nil(t, err)
}
