package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Duration(600 * time.Second),
	}
	response, err := client.Get("http://localhost:8080/close")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("status code %d", response.StatusCode)

}
