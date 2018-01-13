package basic

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
	"unicode"

	"github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

type Options struct {
	Name    string   `url:"name"`
	Address []string `url:"address"`
}

type Rule struct {
	Name string `mapstructure:"rule"`
	Icon string
}

type Options2 struct {
	Width  int `mapstructure:"width"`
	Height int
	Rules  map[string]*Rule
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

func Split(r rune) bool {
	return r == '[' || r == ']'
}
func IsDigital(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
func TestQueryStringEncode3(t *testing.T) {
	raw := "width=11&rules[0][rule]=rule1&rules[0][icon]=icon1&height=123"
	values, err := url.ParseQuery(raw)
	input2 := make(map[string]interface{})
	for key, value := range values {
		key = strings.ToLower(key)
		parts := strings.FieldsFunc(key, Split)
		if len(parts) == 1 {
			input2[parts[0]] = value[0]
		} else if len(parts) == 3 {
			key, index, property := parts[0], parts[1], parts[2]
			if _, exists := input2[key]; !exists {
				input2[key] = make(map[string]interface{})
			}
			container := input2[key].(map[string]interface{})
			if _, exists := container[index]; !exists {
				container[index] = make(map[string]interface{})
			}
			obj := container[index].(map[string]interface{})
			obj[property] = value[0]
		}
	}
	result2 := new(Options2)
	err = mapstructure.WeakDecode(input2, result2)
	assert.Nil(t, err)
	fmt.Printf("query %v\n", values)

	input := map[string]interface{}{
		"width":  "512",
		"height": "256",
		"rules": map[string]map[string]interface{}{
			"0": {
				"rule": "rule1",
				"icon": "icon1",
			},
			"1": {
				"rule": "rule2",
				"icon": "icon2",
			},
		},
	}

	result := new(Options2)
	err = mapstructure.WeakDecode(input, result)
	assert.Nil(t, err)
	fmt.Printf("result :%v", result)
}
