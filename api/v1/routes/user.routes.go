package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joehann-9s/gtd-api/api/v1/handlers"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/user")

	user.Get("/", handlers.GetAllUsers)

	// user.Post("/")

	// user.Get("/:id")

	// user.Put("/:id")

	// user.Delete("/:id")
}
