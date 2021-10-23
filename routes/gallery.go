package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sutantodadang/adopt-me/v1/handler"
	"github.com/sutantodadang/adopt-me/v1/middleware"
)

type galleryRoute struct {
	galleryHandler handler.GalleryHandler
	authMiddle     middleware.MiddlewareService
}

func NewGalleryRoute(galleryHandler *handler.GalleryHandler, authMiddle middleware.MiddlewareService) *galleryRoute {
	return &galleryRoute{*galleryHandler, authMiddle}
}

func (r *galleryRoute) GalleryRouter(app *fiber.App) {
	route := app.Group("/api/v1/gallery")

	route.Post("/", r.authMiddle.AuthMiddle(), r.galleryHandler.CreateGalleryHandler)
}
