package image

import (
	"testing"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/stretchr/testify/assert"
	"golang.org/x/image/tiff"
	"os"
)

func TestImageServe(t *testing.T) {
	img, err := imgio.Open("1.jpg")
	assert.Nil(t, err)
	assert.NotNil(t, img)
	f, err := os.Create("1.tiff")
	assert.Nil(t, err)
	defer f.Close()

	err = tiff.Encode(f, img, &tiff.Options{})
	assert.Nil(t, err)
}
