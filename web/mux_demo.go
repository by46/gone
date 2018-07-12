package web

import (
	"io/ioutil"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	ioutil.ReadAll(request.Body)
	writer.WriteHeader(http.StatusOK)
	writer.WriteHeader(http.StatusInternalServerError)
}

func MissingContentLength(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	for i := 0; i < 1000; i++ {

		writer.Write([]byte("hello worldhello worldhello worldhello worldhello worldhello worldhello world"))
	}
}
func MuxServe() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handler)
	mux.HandleFunc("/miss", MissingContentLength)
	return mux
}
