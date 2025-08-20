package controllers

import (
	"log"

	"douyin_app/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UploadVideo(c *fiber.Ctx) error {
	// Parse form data
	title := c.FormValue("title")
	description := c.FormValue("description")
	userIDStr := c.FormValue("user_id")

	// Validate form data
	if title == "" || userIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title and user ID are required",
		})
	}

	// Convert user ID to ObjectID
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	// Get the uploaded file
	file, err := c.FormFile("video")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Video file is required",
		})
	}

	// Save the file to a destination (this is a simplified example)
	// In a real application, you would upload this to a cloud storage service
	filename := "uploads/" + file.Filename
	if err := c.SaveFile(file, filename); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save video file",
		})
	}

	// Create a new video
	video := models.Video{
		Title:       title,
		Description: description,
		URL:         filename,
		UserID:      userID,
	}

	// Validate the video
	if err := video.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Save the video to the database
	if err := video.Save(); err != nil {
		log.Printf("Error saving video: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save video",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(video)
}

func ListVideos(c *fiber.Ctx) error {
	// TODO: Implement list videos with pagination
	// For now, return a simple message
	return c.SendString("List videos endpoint")
}

func GetVideo(c *fiber.Ctx) error {
	// Get video ID from params
	videoIDStr := c.Params("id")
	if videoIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Video ID is required",
		})
	}

	// Convert video ID to ObjectID
	videoID, err := primitive.ObjectIDFromHex(videoIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid video ID",
		})
	}

	// Find video in database
	// This is a placeholder implementation
	// In a real application, you would query the database
	video := models.Video{
		ID:    videoID,
		Title: "Sample Video",
		URL:   "http://example.com/video.mp4",
	}

	return c.JSON(video)
}

func LikeVideo(c *fiber.Ctx) error {
	// Get video ID from params
	videoIDStr := c.Params("id")
	if videoIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Video ID is required",
		})
	}

	// Convert video ID to ObjectID
	_, err := primitive.ObjectIDFromHex(videoIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid video ID",
		})
	}

	// Get user ID from context (assuming it's set by auth middleware)
	userClaims := c.Locals("user")
	if userClaims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// In a real application, you would update the video's likes in the database
	// This is a placeholder implementation

	return c.JSON(fiber.Map{
		"message": "Video liked successfully",
	})
}

func CommentVideo(c *fiber.Ctx) error {
	// Get video ID from params
	videoIDStr := c.Params("id")
	if videoIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Video ID is required",
		})
	}

	// Convert video ID to ObjectID
	_, err := primitive.ObjectIDFromHex(videoIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid video ID",
		})
	}

	// Get user ID from context (assuming it's set by auth middleware)
	userClaims := c.Locals("user")
	if userClaims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Parse comment content from request body
	var requestBody struct {
		Content string `json:"content"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// In a real application, you would add the comment to the video in the database
	// This is a placeholder implementation

	return c.JSON(fiber.Map{
		"message": "Comment added successfully",
	})
}