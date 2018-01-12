package basic

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMimeTypeDetect(t *testing.T) {
	file, err := os.Open("../image/1.jpg")
	assert.Nil(t, err)
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	assert.Nil(t, err)

	file.Seek(0, 0)
	contentType := http.DetectContentType(buffer)
	fmt.Printf("%v", contentType)
}

func TestMimeType2(t *testing.T) {
	buf := make([]byte, 0)
	fmt.Printf("buf: %s %d", buf, len(buf))
	buf2 := buf[:0]
	fmt.Printf("buf2: %s", buf2)
}

func TestMimeType3(t *testing.T) {
	extensions, err := mime.ExtensionsByType("image/jpeg")
	assert.Nil(t, err)
	fmt.Printf("extensions : %v", extensions)
}
