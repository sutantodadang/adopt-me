package services

import (
	"context"

	"github.com/sutantodadang/adopt-me/v1/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceCat interface {
	CreateCat(input models.Cat) error
}

type serviceCat struct {
	db *mongo.Client
}

func NewServiceCat(db *mongo.Client) *serviceCat {
	return &serviceCat{db}
}

func (s *serviceCat) CreateCat(input models.Cat) error {
	db := s.db.Database("adopt-me-api").Collection("cats")

	_, err := db.InsertOne(context.Background(), &input)

	if err != nil {
		return err
	}

	return nil
}
