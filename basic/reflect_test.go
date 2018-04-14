package basic

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Stone struct {
	Name string `xml:"name,attr", default:"12322"`
	Age  int
}

func TestReflect(t *testing.T) {
	stone := &Stone{}
	s := reflect.ValueOf(stone).Elem()
	typeOfStone := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%s %s %v %v %v\n", typeOfStone.Field(i).Name,
			f.Type(), f.Interface(), typeOfStone.Field(i).Tag.Get("default"),
			f.Kind())
	}
}

func TestXmlUnmarshal(t *testing.T) {
	content := `<?xml version="1.0" ?><coverage name="123"></coverage>`
	stone := &Stone{}
	err := xml.Unmarshal([]byte(content), &stone)
	assert.Nil(t, err)
	fmt.Printf("name %s", stone.Name)
}
