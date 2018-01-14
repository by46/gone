package third

import (
	"testing"
	"github.com/lazywei/go-opencv/opencv"
)

func TestFaceDetect(t *testing.T) {
	image := opencv.LoadImage("/lena.jpg")

	cascade := opencv.LoadHaarClassifierCascade("haarcascade_frontalface_alt.xml")
	faces := cascade.DetectObjects(image)

	for _, value := range faces {
		opencv.Rectangle(image,
			opencv.Point{value.X() + value.Width(), value.Y()},
			opencv.Point{value.X(), value.Y() + value.Height()},
			opencv.ScalarAll(255.0), 1, 1, 0)
	}

	win := opencv.NewWindow("Face Detection")
	win.ShowImage(image)
	opencv.WaitKey(0)
}
