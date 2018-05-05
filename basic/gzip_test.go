package basic

import (
	"archive/tar"
	"bytes"
	"compress/bzip2"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/filetype.v1"
)

func TestUnGzip(t *testing.T) {

	content, err := ioutil.ReadFile("../tmp.tar.gz")
	assert.Nil(t, err)
	mimeType := http.DetectContentType(content)
	fmt.Printf("mime type: %s\n", mimeType)

	kind, _ := filetype.Match(content)
	fmt.Printf("File type: %s. MIME %s", kind.Extension, kind.MIME.Value)
	gzipReader, err := gzip.NewReader(bytes.NewBuffer(content))
	assert.Nil(t, err)
	defer gzipReader.Close()

	reader := tar.NewReader(gzipReader)

	for {
		header, err := reader.Next()
		switch {
		case err == io.EOF:
			fmt.Print("EOF")
			return
		case err != nil:
			fmt.Print("Error")
			return
		case header == nil:
			continue
		}

		switch header.Typeflag {
		case tar.TypeReg:
			buf := bytes.NewBuffer(nil)
			io.Copy(buf, reader)
			fmt.Printf("filename: %s\n", header.Name)
			fmt.Print(buf.String())

		case tar.TypeDir:
			fmt.Printf("dir filename %s\n", header.Name)
		}
	}
}

func TestUnTar(t *testing.T) {
	content, err := ioutil.ReadFile("../tmp.tar")
	assert.Nil(t, err)
	mimeType := http.DetectContentType(content)
	fmt.Printf("mime type: %s\n", mimeType)

	kind, _ := filetype.Match(content)
	fmt.Printf("File type: %s. MIME %s\n", kind.Extension, kind.MIME.Value)

	buf := bytes.NewBuffer(content)
	reader := tar.NewReader(buf)

	for {
		header, err := reader.Next()
		switch {
		case err == io.EOF:
			fmt.Print("EOF")
			return
		case err != nil:
			fmt.Print("Error")
			return
		case header == nil:
			continue
		}

		switch header.Typeflag {
		case tar.TypeReg:
			buf := bytes.NewBuffer(nil)
			io.Copy(buf, reader)
			fmt.Printf("filename: %s\n", header.Name)
			fmt.Print(buf.String())

		case tar.TypeDir:
			fmt.Printf("dir filename %s\n", header.Name)
		}
	}

}

func TestUnBzip(t *testing.T) {
	content, err := ioutil.ReadFile("../im.tar.bz2")
	assert.Nil(t, err)
	mimeType := http.DetectContentType(content)
	fmt.Printf("mime type: %s\n", mimeType)

	kind, _ := filetype.Match(content)
	fmt.Printf("extension %s  MIME %s\n", kind.Extension, kind.MIME.Value)
	bzipReader := bzip2.NewReader(bytes.NewBuffer(content))

	reader := tar.NewReader(bzipReader)

	for {
		header, err := reader.Next()
		switch {
		case err == io.EOF:
			fmt.Print("EOF")
			return
		case err != nil:
			fmt.Print("Error")
			return
		case header == nil:
			continue
		}

		switch header.Typeflag {
		case tar.TypeReg:
			buf := bytes.NewBuffer(nil)
			io.Copy(buf, reader)
			fmt.Printf("filename: %s %d\n", header.Name, buf.Len())

		case tar.TypeDir:
			fmt.Printf("dir filename %s\n", header.Name)
		}
	}
}
