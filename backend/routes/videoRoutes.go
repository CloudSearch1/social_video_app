package routes

import (
	"douyin_app/controllers"
	"douyin_app/middleware"
	"github.com/gofiber/fiber/v2"
)

func VideoRoutes(app *fiber.App) {
	// Public routes
	app.Get("/api/videos/list", controllers.ListVideos)
	app.Get("/api/videos/:id", controllers.GetVideo)

	// Protected routes
	protected := app.Group("/api/videos", middleware.AuthRequired)
	protected.Post("/upload", controllers.UploadVideo)
	protected.Post("/:id/like", controllers.LikeVideo)
	protected.Post("/:id/comment", controllers.CommentVideo)
}