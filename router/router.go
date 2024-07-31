package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	noteRoutes "github.com/duddyV/notes-api-go-fiber/internal/routes/note"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New()) // Group endpoints with param 'api' and log whenever this endpoint is hit.

	// Setup the Node Routes
	noteRoutes.SetupNoteRoutes(api)

	// Define a root route
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Welcome to the Notes API")
    })
}