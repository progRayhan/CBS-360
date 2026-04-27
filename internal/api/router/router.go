package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/progRayhan/CBS-360/internal/api/handler"
)

func Setup(app *fiber.App, userHandler *handler.UserHandler) {
	api := app.Group("/api/v1")

	api.Get("/users/:id", userHandler.GetUser)
}
