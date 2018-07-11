package third

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"testing"

	"github.com/disintegration/imaging"
	"github.com/esimov/pigo/core"
	"github.com/stretchr/testify/assert"
)

func TestPigo(t *testing.T) {
	cascadeFile, err := ioutil.ReadFile("facefinder")
	if err != nil {
		log.Fatalf("Error reading the cascade file: %v", err)
	}

	src, err := pigo.GetImage("../iceland.jpg")
	if err != nil {
		log.Fatalf("Cannot open the image file: %v", err)
	}

	pixels := pigo.RgbToGrayscale(src)
	cols, rows := src.Bounds().Max.X, src.Bounds().Max.Y

	cParams := pigo.CascadeParams{
		MinSize:     20,
		MaxSize:     200,
		ShiftFactor: 0.1,
		ScaleFactor: 1.1,
	}
	imgParams := pigo.ImageParams{
		Pixels: pixels,
		Rows:   rows,
		Cols:   cols,
		Dim:    cols,
	}

	pigo := pigo.NewPigo()
	// Unpack the binary file. This will return the number of cascade trees,
	// the tree depth, the threshold and the prediction from tree's leaf nodes.
	classifier, err := pigo.Unpack(cascadeFile)
	if err != nil {
		log.Fatalf("Error reading the cascade file: %s", err)
	}

	// Run the classifier over the obtained leaf nodes and return the detection results.
	// The result contains quadruplets representing the row, column, scale and detection score.
	dets := classifier.RunCascade(imgParams, cParams)

	// Calculate the intersection over union (IoU) of two clusters.
	dets = classifier.ClusterDetections(dets, 0.2)
	fmt.Printf("%v", dets)

	src2, err := imaging.Open("../iceland.jpg")
	assert.Nil(t, err)

	for _, det := range dets {
		if det.Q < 5.0 {
			continue
		}
		dst := imaging.New(det.Scale, det.Scale, color.Black)
		src2 = imaging.Paste(src2, dst, image.Pt(det.Col-det.Scale/2, det.Row-det.Scale/2))
	}

	imaging.Save(src2, "black.jpg")

}
