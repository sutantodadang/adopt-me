package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" validate:"required" bson:"name" form:"name"`
	Gender    string             `json:"gender" validate:"required" bson:"gender" form:"gender"`
	Place     string             `json:"place" validate:"required" bson:"place" form:"place"`
	Email     string             `json:"email" validate:"required,email" bson:"email" form:"email"`
	Avatar    string             `json:"avatar,omitempty" bson:"avatar,omitempty" form:"avatar" validate:"omitempty"`
	Phone     string             `json:"phone" validate:"required" bson:"phone" form:"phone"`
	Password  string             `json:"password" validate:"required" bson:"password" form:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email" bson:"email"`
	Password string `json:"password" validate:"required" bson:"password"`
}
