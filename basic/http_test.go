package basic

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptrace"
	"sync"
	"testing"

	"fmt"
	"time"

	"github.com/stretchr/testify/assert"
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

type transport struct {
	current           *http.Request
	connCheckpoint    int64
	connectCheckpoint int64
	dnsCheckpoint     int64
}

var (
	MyDefaultTransport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          200,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
)

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.current = req
	return MyDefaultTransport.RoundTrip(req)
}
func (t *transport) GetConn(host string) {
	t.connCheckpoint = time.Now().UnixNano()
}
func (t *transport) GotConn(info httptrace.GotConnInfo) {
	elapse := time.Now().UnixNano() - t.connCheckpoint
	fmt.Printf("Connection: %v\n", info.Conn.LocalAddr())
	fmt.Printf("Url %s, Got conn: %+v, elapse: %d us\n", t.current.UserAgent(), info, elapse/1000)
}

func (t *transport) ConnectStart(network, addr string) {
	t.connectCheckpoint = time.Now().UnixNano()
}

func (t *transport) ConnectDone(network, addr string, err error) {
	elapse := time.Now().UnixNano() - t.connectCheckpoint
	fmt.Printf("network: %s, addr: %s, error: %v, elapse: %d us\n", network, addr, err, elapse/1000)
}
func (t *transport) DNSStart(info httptrace.DNSStartInfo) {
	t.dnsCheckpoint = time.Now().UnixNano()
}
func (t *transport) DNSDone(info httptrace.DNSDoneInfo) {
	elapse := time.Now().UnixNano() - t.dnsCheckpoint
	fmt.Printf("dns %v, elapse: %d us\n", info, elapse/1000)
}

func TestHttpTraceClient(t *testing.T) {
	trans := &transport{}
	client := http.Client{
		Transport: trans,
	}
	trace := &httptrace.ClientTrace{
		DNSStart:     trans.DNSStart,
		DNSDone:      trans.DNSDone,
		GetConn:      trans.GetConn,
		GotConn:      trans.GotConn,
		ConnectStart: trans.ConnectStart,
		ConnectDone:  trans.ConnectDone,
		PutIdleConn: func(err error) {
			fmt.Printf("Put Idle Connection: %v \n", err)
		},
	}
	var wg sync.WaitGroup

	for j := 0; j <= 50; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i <= 100; i++ {
				req, _ := http.NewRequest("GET", "http://scmesos04/benjamin/test/diagnosis.txt", nil)
				req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
				response, err := client.Do(req)
				assert.Nil(t, err)
				_, err = ioutil.ReadAll(response.Body)
				assert.Nil(t, err)
				response.Body.Close()
				fmt.Printf("response code: %d\n\n", response.StatusCode)
				time.Sleep(100 * time.Millisecond)
			}
		}()
	}

	wg.Wait()
}

func TestHttpSimpleRequest(t *testing.T) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	response, err := client.Get("http://scmesos04/benjamin/test/diagnosis.txt")
	assert.Nil(t, err)
	defer response.Body.Close()
	ioutil.ReadAll(response.Body)

	fmt.Printf("status code: %v", response.StatusCode)
}
