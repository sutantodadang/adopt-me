package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/sutantodadang/adopt-me/v1/db"
	"github.com/sutantodadang/adopt-me/v1/handler"
	"github.com/sutantodadang/adopt-me/v1/routes"
	"github.com/sutantodadang/adopt-me/v1/services"
)

func main()  {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed load env", err.Error())
	}

	key := os.Getenv("MONGO_URI")

	Database := db.GetClient(key)
	
	
	// memanggil service
	userService := services.NewService(Database)
	// memanggil handler
	userHandler := handler.NewUserHandler(userService)

	userRoute := routes.NewUserRoute(userHandler)


	route := app.Group("/api/v1")

	userRoute.UserRouter(route)



	app.Listen(":5000")

}