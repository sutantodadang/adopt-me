package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sutantodadang/adopt-me/v1/handler"
)

type userRoute struct {
	userHandler handler.UserHandler
}

func NewUserRoute(userHandler *handler.UserHandler) *userRoute {
	return &userRoute{*userHandler}
}

func (u *userRoute) UserRouter(route fiber.Router) {
	route.Post("/user", u.userHandler.CreateUserHandler)
	route.Post("/user/login", u.userHandler.LoginUserHandler)

}
