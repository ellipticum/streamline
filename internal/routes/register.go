package routes

import (
	"fmt"
	"github.com/ellipticum/streamline/core/router/mux"
	"net/http"
)

func Register(mux *mux.Structure) {
	mux.Add("test", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello, World!")
	})
}
