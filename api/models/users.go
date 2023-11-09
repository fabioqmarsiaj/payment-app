package models

import (
	"time"
)

// User represents an user using the social media
type User struct {
	Name      string    			`bson:"name,omitempty"`
	Email     string    			`bson:"email,omitempty"`
	Password  string    			`bson:"password,omitempty"`
	CreatedAt time.Time 			`bson:"created_at,omitempty"`
}