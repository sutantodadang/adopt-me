package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/sutantodadang/adopt-me/v1/db"
	"github.com/sutantodadang/adopt-me/v1/handler"
	"github.com/sutantodadang/adopt-me/v1/services"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main()  {
	app := fiber.New()
	app.Use(logger.New())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed load env", err.Error())
	}

	key := os.Getenv("MONGO_URI")

	Database := db.GetClient(key)
	
	err = Database.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Failed connect to db", err.Error())
	} else {
		fmt.Println("Success to connect")
	}

	userService := services.NewService(Database)

	userHandler := handler.NewUserHandler(userService)


	route := app.Group("/api/v1")

	route.Post("/user",userHandler.CreateUserHandler)

	app.Listen(":5000")

}