package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sutantodadang/adopt-me/v1/handler"
	"github.com/sutantodadang/adopt-me/v1/middleware"
)

type catRoute struct {
	catHandler handler.CatHandler
	authMiddle middleware.MiddlewareService
}

func NewCatRoute(catHandler *handler.CatHandler, authMiddle middleware.MiddlewareService) *catRoute {
	return &catRoute{*catHandler, authMiddle}
}

func (u *catRoute) CatRouter(app *fiber.App) {
	route := app.Group("/api/v1/cat")

	route.Post("/", u.authMiddle.AuthMiddle(), u.catHandler.CreateCatHandler)
	route.Get("/", u.authMiddle.AuthMiddle(), u.catHandler.FindAllCatHandler)
}
