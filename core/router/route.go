package router

import (
	"log"
	"net/http"
	"regexp"
)

type Route struct {
	Path        string
	Method      string
	Handler     http.HandlerFunc
	Regexp      *regexp.Regexp
	Middlewares []Middleware
}

func (r *Route) Use(middleware Middleware) {
	r.Middlewares = append(r.Middlewares, middleware)

	log.Printf("Middleware added to %s %s, total: %d\n", r.Method, r.Path, len(r.Middlewares))
}
