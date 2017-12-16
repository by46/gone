package opencv

import (
	"testing"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/stretchr/testify/assert"
	"gocv.io/x/gocv"
)

func TestFaceDetection(t *testing.T) {
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	classifier.Load("haar.xml")
	dst, err := imgio.Open("lena.png")
	assert.Nil(t, err)
	assert.NotNil(t, dst)

}
