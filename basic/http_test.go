package basic

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpClient(t *testing.T) {
	client := http.Client{
		Transport: http.DefaultTransport,
	}
	response, err := client.Get("https://www.baidu.com")
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
}
