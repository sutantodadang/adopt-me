package services

import (
	"context"

	"github.com/sutantodadang/adopt-me/v1/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceUser interface {
	CreateUser(data models.User) error
	LoginUser(email string) (models.User, error)
}

type serviceUser struct {
	db *mongo.Client
}

func NewServiceUser(db *mongo.Client) *serviceUser {
	return &serviceUser{db}
}

func (s *serviceUser) CreateUser(data models.User) error {

	dataBase := s.db.Database("adopt-me-api").Collection("users")

	_, err := dataBase.InsertOne(context.Background(), &data)

	if err != nil {
		return err
	}

	return nil
}

func (s *serviceUser) LoginUser(email string) (models.User, error) {
	dataBase := s.db.Database("adopt-me-api").Collection("users")

	var result models.User

	err := dataBase.FindOne(context.Background(), bson.M{"email": email}).Decode(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}
