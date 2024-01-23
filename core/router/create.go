package router

import "net/http"

func Create(method string) func(string, http.HandlerFunc) *Route {
	return func(path string, handler http.HandlerFunc) *Route {
		regexp := CompileRegexp(path)
		route := &Route{Method: method, Path: path, Handler: handler, Regexp: regexp, Middlewares: []Middleware{}}
		Routes = append(Routes, route)
		return route
	}
}
