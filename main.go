package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sutantodadang/adopt-me/v1/db"
	"github.com/sutantodadang/adopt-me/v1/handler"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main()  {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed load env", err.Error())
	}

	key := os.Getenv("MONGO_URI")

	db := db.GetClient(key)
	
	err = db.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Failed connect to db", err.Error())
	} else {
		fmt.Println("Success to connect")
	}

	app.Get("/",handler.HelloWorld)

	app.Listen(":5000")

	fmt.Println("Adopt ME")
}