package image

import (
	"image/color"
	"testing"

	"github.com/disintegration/imaging"
	"github.com/stretchr/testify/assert"
)

func TestNewImage(t *testing.T) {
	input := imaging.New(256, 256, color.Transparent)
	err := imaging.Save(input, "bg.png")
	assert.Nil(t, err)
}
