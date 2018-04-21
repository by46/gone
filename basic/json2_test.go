package basic

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type site struct {
	Name string `json:"-SITE.NAME"`
}
type sites struct {
	Site []*site `json:"SITE"`
}

type Result struct {
	Cmd *sites `json:"appcmd"`
}

func TestJSON3(t *testing.T) {

	entity := FileInfo{
		InDate: time.Now(),
	}
	fmt.Printf("datetime format %v\n", entity)
}

func TestJSON4(t *testing.T) {
	content := `
{
    "appcmd": {
        "SITE": [
            {
                "-SITE.NAME": "Default Web Site",
                "-SITE.ID": "1",
                "-bindings": "http/*:80:",
                "-state": "Started"
            },
            {
                "-SITE.NAME": "DA\nE",
                "-SITE.ID": "2",
                "-bindings": "http/*:8090:",
                "-state": "Started"
            }
        ]
    }
}
`
	result := new(Result)
	err := json.Unmarshal([]byte(content), result)
	if err != nil {
		fmt.Printf("error %v\n", err)
	}
	fmt.Printf("result %v\n", result.Cmd.Site[0].Name)
}
