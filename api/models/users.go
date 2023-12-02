package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents an user using the social media
type User struct {
	ID 		  primitive.ObjectID	`bson:"id,omitempty"`
	Name      string    			`bson:"name,omitempty"`
	Email     string    			`bson:"email,omitempty"`
	Password  string    			`bson:"password,omitempty"`
	CreatedAt time.Time 			`bson:"created_at,omitempty"`
}