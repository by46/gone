package third

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"gopkg.in/h2non/filetype.v1"
)

func TestDetectFile(t *testing.T) {
	//buf, _ := ioutil.ReadFile("tmp.tar.gz")
	buf, _ := ioutil.ReadFile("C:\\Users\\by46\\Desktop\\im.tar")

	prev := time.Now()
	for i := 0; i < 1; i++ {
		kind, unknown := filetype.Match(buf)
		if unknown != nil {
			fmt.Printf("Unknown: %s", unknown)
			return
		}

		fmt.Printf("File type: %s. MIME: %s\n", kind.Extension, kind.MIME.Value)
	}
	fmt.Printf("elapse %v", time.Now().Sub(prev))
}
