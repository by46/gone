package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Duration(20 * time.Second),
	}
	response, err := client.Get("http://localhost:8080/home")
	fmt.Printf("hello %v, %v", response.StatusCode, err)
}
