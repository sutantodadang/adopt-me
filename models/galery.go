package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Gallery struct {
	Id      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Cat_Id  string             `json:"cat_id" bson:"cat_id"`
	User_id string             `json:"user_id" bson:"user_id"`
	Images  []Image            `json:"images" bson:"images"`
}

type Image struct {
	Id          string `json:"id,omitempty" bson:"id"`
	Filename    string `json:"filename" bson:"filename"`
	Image_url   string `json:"image_url" bson:"image_url"`
	Display_url string `json:"display_url" bson:"display_url"`
	Thumb       string `json:"thumb" bson:"thumb"`
	Mime        string `json:"mime" bson:"mime"`
	Extension   string `json:"extension" bson:"extension"`
	Delete_url  string `json:"delete_url" bson:"delete_url"`
}

type ResponseGallery struct {
	Data ImageData `json:"data"`
}

type ImageData struct {
	Delete_url  string        `json:"delete_url"`
	Id          string        `json:"id"`
	Image       ImageResponse `json:"image"`
	Display_url string        `json:"display_url"`
	Thumb       ThumbResponse `json:"thumb"`
}

type ImageResponse struct {
	Extension string `json:"extension"`
	Filename  string `json:"filename"`
	Mime      string `json:"mime"`
	Url       string `json:"url"`
}

type ThumbResponse struct {
	Url string `json:"url"`
}
