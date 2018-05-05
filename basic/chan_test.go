package basic

import (
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	var writeCh <-chan time.Time
	writeCh <- time.Now()
}
