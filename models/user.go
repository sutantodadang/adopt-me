package models

import "time"


type User struct {
	Name string 	`json:"name" validate:"required" bson:"name"`
	Gender string 	`json:"gender" validate:"required" bson:"gender"`
	Place string 	`json:"place" validate:"required" bson:"place"`
	Email string 	`json:"email" validate:"required,email" bson:"email"`
	Avatar string 	`json:"avatar" bson:"avatar"`
	Phone string 	`json:"phone" validate:"required" bson:"phone"`
	Password string `json:"password" validate:"required" bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

