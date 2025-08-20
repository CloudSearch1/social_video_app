package models

import (
	"context"
	"log"
	"time"

	"douyin_app/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const VideoCollection = "videos"

type Video struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	URL         string             `json:"url" bson:"url"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	Likes       []string           `json:"likes" bson:"likes"`
	Comments    []Comment          `json:"comments" bson:"comments"`
	ViewCount   int                `json:"view_count" bson:"view_count"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type Comment struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

func (v *Video) Save() error {
	collection := config.DB.Collection(VideoCollection)

	// Set timestamps
	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()

	// Insert the video into the database
	_, err := collection.InsertOne(context.TODO(), v)
	if err != nil {
		log.Printf("Error inserting video: %v", err)
		return err
	}

	return nil
}

func (v *Video) Validate() error {
	// Check if title is empty
	if v.Title == "" {
		return &ValidationError{Field: "title", Message: "Title is required"}
	}

	// Check if URL is empty
	if v.URL == "" {
		return &ValidationError{Field: "url", Message: "URL is required"}
	}

	// Check if UserID is empty
	if v.UserID == primitive.NilObjectID {
		return &ValidationError{Field: "user_id", Message: "User ID is required"}
	}

	return nil
}