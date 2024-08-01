package userRoutes

import (
    "github.com/gofiber/fiber/v2"
    userHandler "github.com/duddyV/notes-api-go-fiber/internal/handlers/user"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")
    // Authentication routes
    user.Post("/register", userHandler.Register)
    user.Post("/login", userHandler.Login)
}
