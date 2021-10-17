package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sutantodadang/adopt-me/v1/models"
	"github.com/sutantodadang/adopt-me/v1/services"
	"github.com/sutantodadang/adopt-me/v1/utils"
)

type CatHandler struct {
	catService services.ServiceCat
}

func NewCatHandler(catService services.ServiceCat) *CatHandler {
	return &CatHandler{catService}
}

func (h *CatHandler) CreateCatHandler(c *fiber.Ctx) error {
	cat := new(models.Cat)

	user := c.Locals("user").(*jwt.Token)
	claim := user.Claims.(jwt.MapClaims)
	user_id := claim["id"].(string)

	if err := c.BodyParser(cat); err != nil {
		return c.Status(fiber.ErrUnprocessableEntity.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := utils.ValidateInput(cat); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	cat.CreatedAt = time.Now()
	cat.UpdatedAt = time.Now()
	cat.UserId = user_id

	if err := h.catService.CreateCat(*cat); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "successfully created",
	})
}

func (h *CatHandler) FindAllCatHandler(c *fiber.Ctx) error {
	query := c.Query("user_id")

	if query == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "fill query",
		})
	}

	res, err := h.catService.FindCatById(query)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "retrieve data successfully",
		"data":    res,
	})

}