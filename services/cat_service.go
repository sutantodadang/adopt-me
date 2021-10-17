package services

import (
	"context"

	"github.com/sutantodadang/adopt-me/v1/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceCat interface {
	CreateCat(input models.Cat) error
	FindCatById(id string) ([]models.Cat, error)
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

func (s *serviceCat) FindCatById(id string) ([]models.Cat, error) {
	db := s.db.Database("adopt-me-api").Collection("cats")

	var result []models.Cat

	filterOption := options.Find()
	filterOption.SetLimit(10)

	res, err := db.Find(context.Background(), bson.M{"user_id": id}, filterOption)

	if err != nil {

		return nil, err
	}

	if err := res.All(context.Background(), &result); err != nil {

		return nil, err
	}

	if result == nil {
		return nil, mongo.ErrNoDocuments
	}

	return result, nil
}
