package basic

import (
	"testing"
	"strings"
	"io"
	"os"
	"github.com/labstack/gommon/log"
	"fmt"
	"bytes"
	"io/ioutil"
)

func TestCopyDemo(t *testing.T) {
	r := strings.NewReader("io. examples")
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func TestCopyBuffer(t *testing.T) {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}
}

func TestCopyN(t *testing.T) {
	r := strings.NewReader("some thing\n")

	if _, err := io.CopyN(os.Stdout, r, 5); err != nil {
		log.Fatal(err)
	}
}

func TestReadAtLeast(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")
	buf := make([]byte, 33)
	if _, err := io.ReadAtLeast(r, buf, 4); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", buf)

	shortBuf := make([]byte, 3)
	if _, err := io.ReadAtLeast(r, shortBuf, 4); err != nil {
		fmt.Printf("error: %s\n", err)
	}
	bigBuf := make([]byte, 64)
	if _, err := io.ReadAtLeast(r, bigBuf, 64); err != nil {
		fmt.Printf("error: %s\n", err)
	}
}

func TestReadFull(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")
	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	buf = make([]byte, 64)
	if _, err := io.ReadFull(r, buf); err != nil {
		fmt.Printf("error: %s\n", err)
	}
	io.WriteString(os.Stdout, "hello world\n")

	r = strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 4)
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}

func TestMultiReader(t *testing.T) {
	r1 := strings.NewReader("hello world\n")
	r2 := strings.NewReader("I am good.\n")
	r3 := strings.NewReader("last reader\n")

	r := io.MultiReader(r1, r2, r3)
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func TestTeeReader(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read")
	var buf bytes.Buffer
	tee := io.TeeReader(r, &buf)
	printall := func(r io.Reader) {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", b)
	}
	printall(tee)
	printall(&buf)

	r = strings.NewReader("some io.Reader stream to be read")
	s := io.NewSectionReader(r, 5, 17)
	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}
}
