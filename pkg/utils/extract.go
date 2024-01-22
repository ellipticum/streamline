package utils

import (
	"net/http"
	"strings"
)

func Extract(r *http.Request) string {
	path := r.URL.Path

	path = strings.TrimSuffix(path, "/")

	parts := strings.Split(path, "/")

	if len(parts) > 0 {
		return parts[len(parts)-1]
	}

	return ""
}
