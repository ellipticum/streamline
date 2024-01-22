package router

import (
	"net/http"
	"regexp"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
	Regexp  *regexp.Regexp
}
