package basic

import (
	"testing"
	"net/url"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestPathEscape(t *testing.T) {
	name := "file name"
	escapeName := url.PathEscape(name)
	assert.Equal(t, "file%20name", escapeName)
	escapeName2 := fmt.Sprintf("%s", escapeName)
	assert.Equal(t, "file%20name", escapeName2)
}
