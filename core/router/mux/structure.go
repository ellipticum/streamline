package mux

import (
	"net/http"
)

type Structure struct {
	routes map[string]http.HandlerFunc
}

func (mux *Structure) Add(path string, handler http.HandlerFunc) {
	mux.routes[path] = handler
}

func (mux *Structure) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if handler, ok := mux.routes[request.URL.Path]; ok {
		handler(writer, request)
	} else {
		http.NotFound(writer, request)
	}
}
