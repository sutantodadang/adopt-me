package models


type User struct {
	Name string `json:"name" binding:"required" bson:"name"`
	Gender string `json:"gender" binding:"required" bson:"gender"`
	Place string `json:"place" binding:"required" bson:"place"`
	Email string `json:"email" binding:"required,email" bson:"email"`
	Avatar string `json:"avatar" bson:"avatar"`
	Phone string `json:"phone" binding:"required" bson:"phone"`
	Password string `json:"password" binding:"required" bson:"password"`

}

