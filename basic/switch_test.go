package basic

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gorilla/http"
)

func TestSwitchNormal(t *testing.T) {
	var b interface{}
	b = fmt.Errorf("ook")
	switch v := b.(type) {
	case *http.StatusError:
		fmt.Printf("ok %v\n", v)
	}
	segments := strings.Split("a_", "_")
	fmt.Printf("%v %d\n", segments, len(segments))
}
