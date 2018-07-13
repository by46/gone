package basic

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDialer(t *testing.T) {
	dialer := net.Dialer{
		KeepAlive: time.Second * 5,
	}
	conn, err := dialer.Dial("tcp", "scmesos04:80")
	assert.Nil(t, err)
	n, err := conn.Write([]byte("hello"))
	assert.Nil(t, err)
	fmt.Printf("write count: %v\n", n)
	conn.Close()
}

func TestDialerWithTimeout(t *testing.T) {
	dialer := net.Dialer{
		Deadline: time.Now().Add(time.Second * 20),
	}

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	conn, err := dialer.DialContext(ctx, "tcp", "scmesos04:08")
	assert.Nil(t, err)
	conn.Close()
}
