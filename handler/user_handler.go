package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/sutantodadang/adopt-me/v1/helpers"
	"github.com/sutantodadang/adopt-me/v1/models"
	"github.com/sutantodadang/adopt-me/v1/services"
	"github.com/sutantodadang/adopt-me/v1/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userService services.ServiceUser
	jwtService  utils.JwtToken
}

func NewUserHandler(userService services.ServiceUser, jwt utils.JwtToken) *UserHandler {
	return &UserHandler{userService, jwt}
}

func (h *UserHandler) CreateUserHandler(c *fiber.Ctx) error {

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})

	}

	// validasi input
	if err := utils.ValidateInput(user); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	//  bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	user.Password = string(hash)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = h.userService.CreateUser(*user)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "successfully created"})
}

func (h *UserHandler) LoginUserHandler(c *fiber.Ctx) error {
	user := new(models.UserLogin)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	if err := utils.ValidateInput(user); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	res, err := h.userService.LoginUser(user.Email)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(user.Password))

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "incorrect password, try again"})
	}

	token, err := h.jwtService.GenerateToken(res)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
	}

	response := helpers.ResponseApi("success", fiber.Map{"token": token})

	return c.Status(fiber.StatusOK).JSON(response)

}
