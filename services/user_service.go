package services

import (
	"context"

	"github.com/sutantodadang/adopt-me/v1/models"
	"go.mongodb.org/mongo-driver/mongo"
)


type ServiceUser interface {
	CreateUser(data models.User) error
}

type service struct {
	db *mongo.Client
}

func NewService(db *mongo.Client) *service  {
	return &service{db}
}

func (s *service) CreateUser( data models.User) error {
	
	dataBase := s.db.Database("adopt-me-api").Collection("users")

	_, err := dataBase.InsertOne(context.Background(), data)

	if err != nil {
		return err
	}


	return  nil
}