package router

import (
	"regexp"
	"strings"
)

func transformPattern(pattern string) string {
	segments := strings.Split(pattern, "/")
	for i, segment := range segments {
		if strings.HasPrefix(segment, ":") {
			segments[i] = "([^/]+)"
		}
	}
	return "^" + strings.Join(segments, "/") + "$"
}

func CompileRegexp(pattern string) *regexp.Regexp {
	regexPattern := transformPattern(pattern)
	return regexp.MustCompile(regexPattern)
}
