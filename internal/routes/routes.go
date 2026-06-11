package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeon-prince/ainyx-go-backend/internal/handler"
)

// SetupRoutes registers all the API endpoints to their respective handlers.
func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	api := app.Group("/users")

	api.Post("/", userHandler.CreateUser)
	api.Get("/:id", userHandler.GetUser)
	api.Delete("/:id", userHandler.DeleteUser)
	// Update and List endpoints will be registered here
}
