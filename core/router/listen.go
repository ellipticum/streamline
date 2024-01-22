package router

import "net/http"

func Listen(w http.ResponseWriter, r *http.Request) {
	for _, route := range Routes {
		if r.Method == route.Method && route.Regexp.MatchString(r.URL.Path) {
			route.Handler(w, r)
			return
		}
	}
	http.NotFound(w, r)
}
