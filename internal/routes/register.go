package routes

import (
	"github.com/ellipticum/streamline/core/router"
	"github.com/ellipticum/streamline/internal/controllers"
	"github.com/ellipticum/streamline/internal/middlewares"
)

var controller = controllers.User{}

func Register() {
	router.Get("/api/v1/users", controller.Get).Use(middlewares.Logging)
	router.Get("/api/v1/users/:id", controller.GetByID).Use(middlewares.Logging)
	router.Post("/api/v1/users", controller.Create).Use(middlewares.Logging)
}
