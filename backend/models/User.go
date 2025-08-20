package models

import (
	"context"
	"log"
	"time"

	"douyin_app/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const UserCollection = "users"

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func (u *User) Save() error {
	collection := config.DB.Collection(UserCollection)

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	// Set timestamps
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	// Insert the user into the database
	_, err = collection.InsertOne(context.TODO(), u)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return err
	}

	return nil
}

func (u *User) Validate() error {
	// Check if username is empty
	if u.Username == "" {
		return &ValidationError{Field: "username", Message: "Username is required"}
	}

	// Check if email is empty
	if u.Email == "" {
		return &ValidationError{Field: "email", Message: "Email is required"}
	}

	// Check if password is empty
	if u.Password == "" {
		return &ValidationError{Field: "password", Message: "Password is required"}
	}

	// Check if password is at least 6 characters
	if len(u.Password) < 6 {
		return &ValidationError{Field: "password", Message: "Password must be at least 6 characters"}
	}

	return nil
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// FindUserByEmail finds a user by email
func FindUserByEmail(email string) (*User, error) {
	collection := config.DB.Collection(UserCollection)
	var user User
	filter := bson.M{"email": email}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// CheckPassword checks if the provided password matches the hashed password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}