package third

import (
	"fmt"
	"io/ioutil"
	"testing"

	"gopkg.in/h2non/filetype.v1"
)

func TestDetectFile(t *testing.T) {
	buf, _ := ioutil.ReadFile("lena.jpg")
	kind, unknown := filetype.Match(buf)
	if unknown != nil {
		fmt.Printf("Unknown: %s", unknown)
		return
	}

	fmt.Printf("File type: %s. MIME: %s\n", kind.Extension, kind.MIME.Value)
}
