package routes

import (
	"github.com/ellipticum/streamline/core/router"
	"github.com/ellipticum/streamline/internal/controllers"
)

var controller = controllers.User{}

func Register() {
	router.Get("/api/v1/users", controller.Get)
	router.Post("/api/v1/users", controller.Create)
}
