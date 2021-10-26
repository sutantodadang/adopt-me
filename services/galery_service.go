package services

import (
	"context"

	"github.com/sutantodadang/adopt-me/v1/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceGalery interface {
	CreateGallery(gallery models.Gallery) error
	FindGalleryByUserId(id string, limit int) ([]models.Gallery, error)
	FindGalleryByCatId(id string) (models.Gallery, error)
	FindAllGallery(limit int) ([]models.Gallery, error)
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

func (s *serviceGalery) FindGalleryByUserId(id string, limit int) ([]models.Gallery, error) {
	db := s.db.Database("adopt-me-api").Collection("gallery")

	opt := options.Find().SetLimit(int64(limit))

	res, err := db.Find(context.Background(), bson.M{"user_id": id}, opt)

	if err != nil {
		return nil, err
	}

	var result []models.Gallery

	if err := res.All(context.Background(), &result); err != nil {
		return nil, err
	}

	if result == nil {
		return nil, mongo.ErrNoDocuments
	}

	return result, nil

}

func (s *serviceGalery) FindGalleryByCatId(id string) (models.Gallery, error) {
	db := s.db.Database("adopt-me-api").Collection("gallery")

	var result models.Gallery

	if err := db.FindOne(context.Background(), bson.M{"cat_id": id}).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

func (s *serviceGalery) FindAllGallery(limit int) ([]models.Gallery, error) {
	db := s.db.Database("adopt-me-api").Collection("gallery")

	option := options.Find().SetLimit(int64(limit))

	res, err := db.Find(context.Background(), bson.D{}, option)

	if err != nil {
		return nil, err
	}

	var result []models.Gallery

	err = res.All(context.Background(), &result)

	if err != nil {
		return nil, err
	}

	if result == nil {
		return result, nil
	}

	return result, nil

}
