package main

import (
    "github.com/gofiber/fiber/v2"
	"github.com/duddyV/notes-api-go-fiber/database"
	"github.com/duddyV/notes-api-go-fiber/router"
)

func main() {
    // Start a new fiber app
    app := fiber.New()

	database.ConnectDB()

    // Setup the router
	router.SetupRoutes(app)

    // Listen on PORT 3000
    app.Listen(":3000")
}