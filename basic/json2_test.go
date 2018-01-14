package basic

import (
	"fmt"
	"testing"
	"time"
)

func TestJSON3(t *testing.T) {
	entity := FileInfo{
		InDate: time.Now(),
	}
	fmt.Printf("datetime format %v\n", entity)
}
