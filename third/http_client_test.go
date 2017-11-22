package third

import (
	"testing"
	"github.com/gorilla/http"
	"os"
	"github.com/magiconair/properties/assert"
)

func TestClient(t *testing.T) {
	url := "http://localhost:8080/file"
	reader, err := os.Open("http_client_test.go")
	if err != nil {
		assert.Equal(t, err, nil)
	}
	defer reader.Close()

	err = http.Post(url, reader)
	assert.Equal(t, err, nil)
}
