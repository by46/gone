package image

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"os"
	"testing"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/disintegration/imaging"
	"github.com/golang/freetype"
	"github.com/stretchr/testify/assert"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func TestFont(t *testing.T) {
	dst, err := imgio.Open("nasa.jpg")
	assert.Nil(t, err)
	x, y := 10, 10
	point := fixed.Point26_6{
		X: fixed.Int26_6(x * 64),
		Y: fixed.Int26_6(y * 64)}
	img := image.NewRGBA(image.Rect(0, 0, 300, 100))
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.Black),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString("hello world")
	dst = imaging.Overlay(dst, img, image.Pt(0, 0), 1)
	writer, err := os.OpenFile("font-overlay.jpeg", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer writer.Close()
	assert.Nil(t, err)
	err = imaging.Encode(writer, dst, imgio.PNG)
	assert.Nil(t, err)
}

func TestFreeType(t *testing.T) {
	text := "hello world"
	dst, err := imgio.Open("nasa.jpg")
	assert.Nil(t, err)

	img := image.NewRGBA(dst.Bounds())
	fg := image.Black

	fontBytes, err := ioutil.ReadFile("luxisr.ttf")
	assert.Nil(t, err)
	f, err := freetype.ParseFont(fontBytes)
	assert.Nil(t, err)

	fontSize := 40.0
	c := freetype.NewContext()
	c.SetFont(f)
	c.SetFontSize(fontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(fg)
	c.SetHinting(font.HintingFull)

	fmt.Printf("pt")
	pt := freetype.Pt(0, 0+int(c.PointToFixed(fontSize)>>6))
	fmt.Printf("pt %v", pt)
	_, err = c.DrawString(text, pt)
	assert.Nil(t, err)

	dst = imaging.Overlay(dst, img, image.Pt(0, 0), 1)
	writer, err := os.OpenFile("font-overlay.jpeg", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer writer.Close()
	assert.Nil(t, err)
	err = imaging.Encode(writer, dst, imgio.PNG)
	assert.Nil(t, err)
}
