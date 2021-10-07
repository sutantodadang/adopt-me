package handler

import (
	"context"

	"github.com/sutantodadang/adopt-me/v1/models"
	"go.mongodb.org/mongo-driver/mongo"
)


func CreateUser( db *mongo.Client, data models.User) (*mongo.InsertOneResult, error) {
	dataBase := db.Database("adopt-me-api").Collection("users")

	res, err := dataBase.InsertOne(context.Background(), data)

	if err != nil {
		return nil ,err
	}


	return res, nil
}