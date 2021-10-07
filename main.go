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
	"github.com/sutantodadang/adopt-me/v1/models"
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

	route := app.Group("/api/v1")

	route.Post("/user",func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		gender := c.FormValue("gender")
		place := c.FormValue("place")
		email := c.FormValue("email")
		avatar := c.FormValue("avatar")
		phone := c.FormValue("phone")
		password := c.FormValue("password")

		var user models.User

		user.Name = name
		user.Gender = gender
		user.Place = place
		user.Avatar = avatar
		user.Email = email
		user.Phone = phone
		user.Password = password
		
		res,err := handler.CreateUser(Database, user)
		if err != nil {
			c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message":err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(res)
	})

	app.Listen(":5000")

}