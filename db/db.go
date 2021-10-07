package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func GetClient() *mongo.Client  {
	

	clientOptions := options.Client().
    ApplyURI("mongodb+srv://admin:admin@adopt-me.ixrkk.mongodb.net/adopt-me?retryWrites=true&w=majority")
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, clientOptions)
if err != nil {
    log.Fatal(err)
}

return client

}