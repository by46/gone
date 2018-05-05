package basic

import (
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	writeCh := make(chan time.Time)
	writeCh <- time.Now()
	close(writeCh)
	<-writeCh
}
