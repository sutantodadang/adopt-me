package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sutantodadang/adopt-me/v1/models"
	"github.com/sutantodadang/adopt-me/v1/services"
)


type userHandler struct {
	userService services.ServiceUser
}

func NewUserHandler(userService services.ServiceUser) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) CreateUserHandler(c *fiber.Ctx) error {
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

	// kurang bcrypt
	
	err := h.userService.CreateUser(user)
	if err != nil {
		c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message":err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message":"successfully created"})
}