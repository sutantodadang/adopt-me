package services

import (
	"context"

	"github.com/sutantodadang/adopt-me/v1/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceGalery interface {
	CreateGallery(gallery models.Gallery) error
}

type serviceGalery struct {
	db *mongo.Client
}

func NewServiceGalery(db *mongo.Client) *serviceGalery {
	return &serviceGalery{db}
}

func (s *serviceGalery) CreateGallery(gallery models.Gallery) error {

	db := s.db.Database("adopt-me-api").Collection("gallery")

	_, err := db.InsertOne(context.Background(), &gallery)

	if err != nil {
		return err
	}

	return nil
}
