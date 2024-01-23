package router

import (
	"log"
	"net/http"
)

func Listen(w http.ResponseWriter, r *http.Request) {
	for _, route := range Routes {
		if r.Method == route.Method && route.Regexp.MatchString(r.URL.Path) {
			handler := route.Handler

			log.Printf("Checking route: %s, middleware count: %d\n", route.Path, len(route.Middlewares))

			for i := len(route.Middlewares) - 1; i >= 0; i-- {
				handler = route.Middlewares[i](handler)
			}

			handler(w, r)

			return
		}
	}

	http.NotFound(w, r)
}
