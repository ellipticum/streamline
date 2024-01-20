package mux

import "net/http"

func Create() *Structure {
	return &Structure{routes: make(map[string]http.HandlerFunc)}
}
