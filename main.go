package main

import (
    "github.com/gofiber/fiber/v2"
	"github.com/duddyV/notes-api-go-fiber/database"
)

func main() {
    // Start a new fiber app
    app := fiber.New()

	database.ConnectDB()

    // Send a string back for GET calls to the endpoint "/"
    app.Get("/", func(c *fiber.Ctx) error {
        err := c.SendString("And the API is UP! test")
        return err
    })

    // Listen on PORT 3000
    app.Listen(":3000")
}