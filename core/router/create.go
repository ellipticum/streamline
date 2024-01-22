package router

import "net/http"

func Create(method string) func(string, http.HandlerFunc) {
	return func(path string, handler http.HandlerFunc) {
		regexp := CompileRegexp(path)
		Routes = append(Routes, Route{Method: method, Path: path, Handler: handler, Regexp: regexp})
	}
}
