package image

import (
	"testing"

	"github.com/disintegration/imaging"
	"github.com/stretchr/testify/assert"
)

func TestRadius(t *testing.T) {
	img, err := imaging.Open("sample.png")
	assert.Nil(t, err)
	img = Radius1(img, 100)
	err = imaging.Save(img, "lena-radius.png")
	assert.Nil(t, err)
}
