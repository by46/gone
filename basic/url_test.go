package basic

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

type Options struct {
	Name    string   `url:"name"`
	Address []string `url:"address"`
}

func TestPathEscape(t *testing.T) {
	name := "file name"
	escapeName := url.PathEscape(name)
	assert.Equal(t, "file%20name", escapeName)
	escapeName2 := fmt.Sprintf("%s", escapeName)
	assert.Equal(t, "file%20name", escapeName2)
}

func TestQueryStringEncode(t *testing.T) {
	address := "http://localhost/home?rules[0][rule]=rule1&rules[0][icon]=icon1"
	options, err := url.Parse(address)
	assert.Nil(t, err)
	fmt.Printf("query %v", options.Query())
}

func TestQueryStringEncode2(t *testing.T) {
	input := []map[string]interface{}{{
		"Name":    "Mitchell",
		"Address": []string{"one", "two", "three"},
	}}
	person := make([]*Options, 0)
	err := mapstructure.Decode(input, &person)
	assert.Nil(t, err)
	fmt.Printf("mapping %v", person)
}

func TestQueryStringWithObject(t *testing.T) {
	options := &Options{
		Name:    "benjamin",
		Address: []string{"Chengdu", "Chongqing"},
	}
	v, err := query.Values(options)
	assert.Nil(t, err)
	fmt.Printf("query : %v", v.Encode())
}
