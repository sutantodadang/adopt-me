package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/sutantodadang/adopt-me/v1/models"
	"github.com/sutantodadang/adopt-me/v1/services"
	"github.com/sutantodadang/adopt-me/v1/utils"
	"golang.org/x/crypto/bcrypt"
)


type UserHandler struct {
	userService services.ServiceUser
}

func NewUserHandler(userService services.ServiceUser) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) CreateUserHandler(c *fiber.Ctx) error {

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message":err.Error()})
		return err
	}

	// validasi input
	if err := utils.ValidateInput(*user); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message":err.Error()})
	}

	//  bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message":err.Error()})
	}

	user.Password = string(hash)
	
	
	err = h.userService.CreateUser(*user)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message":err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message":"successfully created"})
}