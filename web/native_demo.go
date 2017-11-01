package web

import "net/http"

type NativeHandler struct {
}

func NewNativeHandler() *NativeHandler {
	return &NativeHandler{}
}

func (h *NativeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world."))
}
func NativeServe() http.Handler {
	handler := NewNativeHandler()
	return handler
}
