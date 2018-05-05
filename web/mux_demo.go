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

func MuxServe() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handler)
	return mux
}
