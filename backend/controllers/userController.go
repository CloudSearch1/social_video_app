package controllers

import (
	"log"
	"os"
	"time"

	"douyin_app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RegisterUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate user input
	if err := user.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Check if user already exists
	existingUser, err := models.FindUserByEmail(user.Email)
	if err != nil {
		log.Printf("Error checking existing user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if existingUser != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	// Save user to database
	if err := user.Save(); err != nil {
		log.Printf("Error saving user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to register user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func LoginUser(c *fiber.Ctx) error {
	var requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Find user by email
	user, err := models.FindUserByEmail(requestBody.Email)
	if err != nil {
		log.Printf("Error finding user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Check password
	if !user.CheckPassword(requestBody.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Generate JWT token
	claims := jwt.MapClaims{
		"user_id": user.ID.Hex(),
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not login",
		})
	}

	return c.JSON(fiber.Map{
		"token": t,
	})
}

func GetUserProfile(c *fiber.Ctx) error {
	// Get user ID from context (set by auth middleware)
	userClaims := c.Locals("user")
	if userClaims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// In a real application, you would fetch the user profile from the database
	// This is a placeholder implementation
	return c.JSON(fiber.Map{
		"message": "User profile",
		"user":    userClaims,
	})
}

func UpdateUserProfile(c *fiber.Ctx) error {
	// Get user ID from context (set by auth middleware)
	userClaims := c.Locals("user")
	if userClaims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Parse request body
	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// In a real application, you would update the user profile in the database
	// This is a placeholder implementation
	return c.JSON(fiber.Map{
		"message": "Profile updated successfully",
		"updates": updates,
	})
}