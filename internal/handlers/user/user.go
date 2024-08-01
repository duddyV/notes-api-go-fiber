package userHandler

import (
    "github.com/gofiber/fiber/v2"
    "golang.org/x/crypto/bcrypt"

    "github.com/duddyV/notes-api-go-fiber/database"
    userModel "github.com/duddyV/notes-api-go-fiber/internal/models/user"
)

func Register(c *fiber.Ctx) error {
    user := new(userModel.User)

    if err := c.BodyParser(user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
    }
    user.Password = string(hashedPassword)

    // Save user in the database
    if err := database.DB.Create(&user).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func Login(c *fiber.Ctx) error {
    var input struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    var user userModel.User
    if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    // Compare hashed password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Login successful"})
}
