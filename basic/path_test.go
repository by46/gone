package basic

import (
	"fmt"
	"net/url"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathJoin(t *testing.T) {
	assert.Equal(t, "a/b", path.Join("a", "b"))
	//assert.Equal(t, "http://example/b", path.Join("http://example", "b"))
	u, _ := url.Parse("http://example.org/a/abcd/abcd/b/c")
	fmt.Printf("%s %s\n", u.EscapedPath(), u)
}
