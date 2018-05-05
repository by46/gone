package main

import (
	"time"
)

import (
	"net/http"
	"fmt"
	"net"
	"io/ioutil"
)

func HttpClient() {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          20,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   3,
	}
	client := http.Client{
		Transport: transport,
	}
	for j := 0; j <= 1; j++ {
		go func(worker int) {

			for i := 1; i <= 200; i++ {
				if response, err := client.Get("https://www.baidu.com"); err != nil {
					fmt.Printf("error %s\n", err)
				} else {
					fmt.Printf("worker %d, batch %d code %d\n", worker, i, response.StatusCode)
					ioutil.ReadAll(response.Body)
				}
				time.Sleep(time.Millisecond * 200)
			}
		}(j)
	}
	time.Sleep(50 * time.Second)

}

func main() {
	HttpClient()
}
