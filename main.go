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

	// di aktifkan kalau berjalan di lokal
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Failed load env", err.Error())
	// }

	key := os.Getenv("MONGO_URI")
	secret := os.Getenv("SECRET_KEY")
	port := os.Getenv("PORT")
	key_img := os.Getenv("KEY_IMG")
	url_img := os.Getenv("IMG_URL")

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	Database := db.GetClient(key)

	// memanggil service
	userService := services.NewServiceUser(Database)
	catService := services.NewServiceCat(Database)
	galleryService := services.NewServiceGalery(Database)
	jwtService := utils.NewJwt(secret)
	middleService := middleware.NewAuthMiddleware(secret)

	// memanggil handler
	userHandler := handler.NewUserHandler(userService, jwtService)
	catHandler := handler.NewCatHandler(catService)
	galleryHandler := handler.NewGalleryHandler(galleryService, catService, url_img, key_img)

	// memanggil route
	userRoute := routes.NewUserRoute(userHandler, middleService)
	catRoute := routes.NewCatRoute(catHandler, middleService)
	galleryRoute := routes.NewGalleryRoute(galleryHandler, middleService)

	// set route
	userRoute.UserRouter(app)
	catRoute.CatRouter(app)
	galleryRoute.GalleryRouter(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome To Adopt Me Api")
	})

	if port == "" {
		port = "27017"
	}

	app.Listen(":" + port)

}
