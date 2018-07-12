package basic

import (
	"fmt"
	"testing"
	"time"

	"github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

const (
	RFC1123 = "Mon, 02 Jan 2006 15:04:05 GMT"
)

type FileInfo struct {
	InDate time.Time `json:"InDate"`
}

func TestISO8601DateJson(t *testing.T) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	buf := []byte("{\"InDate\":\"2017-11-23T08:00:55Z\"}")
	entity := new(FileInfo)
	err := json.Unmarshal(buf, entity)
	assert.Nil(t, err)
	fmt.Printf("hello %v\n", entity.InDate)
	b := make([]byte, 0, len(RFC1123)+2)
	b = entity.InDate.AppendFormat(b, RFC1123)
	fmt.Printf("datetime format %s\n", b)
}

func TestJSON2(t *testing.T) {
	entity := FileInfo{
		InDate: time.Now(),
	}
	fmt.Printf("datetime format %v\n", entity)
}

func TestJsonUnmarshal(t *testing.T) {
	assertion := assert.New(t)
	content := []byte(`{
		"firstName": "Jean",
		"lastName": "Bartik",
		"age": 86,
		"education": [
		{
		"institution": "Northwest Missouri State Teachers College",
		"degree": "Bachelor of Science in Mathematics"
		},
		{
		"institution": "University of Pennsylvania",
		"degree": "Masters in English"
		}
		],
		"spouse": "William Bartik",
		"children": [
		"Timothy John Bartik",
		"Jane Helen Bartik",
		"Mary Ruth Bartik"
		]
		}`)

	var f interface{}
	err := jsoniter.Unmarshal(content, &f)
	assertion.Nil(err)

	fmt.Println(f)

}
