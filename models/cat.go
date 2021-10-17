package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cat struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" validate:"required" bson:"name"`
	Ras         string             `json:"ras" validate:"required" bson:"ras"`
	Description string             `json:"description" validate:"required" bson:"description"`
	Medical     string             `json:"medical" validate:"required"  bson:"medical"`
	Gender      string             `json:"gender" validate:"required" bson:"gender"`
	Weight      int                `json:"weight" validate:"required" bson:"weight"`
	Height      int                `json:"height" validate:"required" bson:"height"`
	UserId      string             `json:"user_id" bson:"user_id"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}
