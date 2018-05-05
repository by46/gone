package basic

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
	"time"
)

func TestHttpClient(t *testing.T) {
	client := http.Client{
		Transport: http.DefaultTransport,
	}
	for i := 1; i <= 20; i++ {
		response, _ := client.Get("https://www.baidu.com")
		fmt.Printf("worker %d, code %s", i, response.StatusCode)
		time.Sleep(time.Second)
	}
	response, err := client.Get("https://www.baidu.com")
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
}
