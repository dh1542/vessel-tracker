package api

import (
	"net/http"
)

type homeHandler struct{}

func (h homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func ServeHTTPServer() {
	mux := http.NewServeMux()
	mux.Handle("/", &homeHandler{})
	http.ListenAndServe(":8080", mux)
}
