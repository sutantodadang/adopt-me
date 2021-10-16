package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sutantodadang/adopt-me/v1/db"
	"github.com/sutantodadang/adopt-me/v1/handler"
	"github.com/sutantodadang/adopt-me/v1/middleware"
	"github.com/sutantodadang/adopt-me/v1/routes"
	"github.com/sutantodadang/adopt-me/v1/services"
	"github.com/sutantodadang/adopt-me/v1/utils"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// di aktifkan kalau berjalan di lokal
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Failed load env", err.Error())
	// }

	key := os.Getenv("MONGO_URI")
	secret := os.Getenv("SECRET_KEY")

	Database := db.GetClient(key)

	// memanggil service
	userService := services.NewService(Database)
	jwtService := utils.NewJwt(secret)
	middleService := middleware.NewAuthMiddleware(secret)

	// memanggil handler
	userHandler := handler.NewUserHandler(userService, jwtService)

	// memanggil route
	userRoute := routes.NewUserRoute(userHandler, middleService)

	route := app.Group("/api/v1")

	userRoute.UserRouter(route)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome To Adopt Me Api")
	})

	app.Listen(":5050")

}
