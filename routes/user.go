package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sutantodadang/adopt-me/v1/handler"
	"github.com/sutantodadang/adopt-me/v1/middleware"
)

type userRoute struct {
	userHandler handler.UserHandler
	authMiddle  middleware.MiddlewareService
}

func NewUserRoute(userHandler *handler.UserHandler, authMiddle middleware.MiddlewareService) *userRoute {
	return &userRoute{*userHandler, authMiddle}
}

func (u *userRoute) UserRouter(route fiber.Router) {
	route.Post("/user", u.userHandler.CreateUserHandler)
	route.Post("/user/login", u.userHandler.LoginUserHandler)

}
