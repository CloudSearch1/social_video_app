package routes

import (
	"douyin_app/controllers"
	"douyin_app/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	// Public routes
	app.Post("/api/users/register", controllers.RegisterUser)
	app.Post("/api/users/login", controllers.LoginUser)

	// Protected routes
	protected := app.Group("/api/users", middleware.AuthRequired)
	protected.Get("/profile", controllers.GetUserProfile)
	protected.Put("/profile", controllers.UpdateUserProfile)
}