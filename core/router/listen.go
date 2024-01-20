package router

import "net/http"

func Listen(w http.ResponseWriter, r *http.Request) {
	for _, route := range Routes {
		if r.URL.Path == route.Path && r.Method == route.Method {
			route.Handler(w, r)

			return
		}
	}
}
