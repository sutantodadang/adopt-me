package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" validate:"required" bson:"name"`
	Gender    string             `json:"gender" validate:"required" bson:"gender"`
	Place     string             `json:"place" validate:"required" bson:"place"`
	Email     string             `json:"email" validate:"required,email,unique" bson:"email"`
	Avatar    string             `json:"avatar" bson:"avatar"`
	Phone     string             `json:"phone" validate:"required,unique" bson:"phone"`
	Password  string             `json:"password" validate:"required" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email" bson:"email"`
	Password string `json:"password" validate:"required" bson:"password"`
}
