package handler

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sutantodadang/adopt-me/v1/helpers"
	"github.com/sutantodadang/adopt-me/v1/models"
	"github.com/sutantodadang/adopt-me/v1/services"
	"github.com/sutantodadang/adopt-me/v1/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (h *CatHandler) FindAllCatByUserIdHandler(c *fiber.Ctx) error {
	query := c.Query("user_id")

	if query == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "fill query",
		})
	}

	res, err := h.catService.FindCatByUserId(query)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := helpers.ResponseApi("retrieve data successfully", res)

	return c.Status(fiber.StatusOK).JSON(response)

}

func (h *CatHandler) FindCatHandler(c *fiber.Ctx) error {
	id := c.Query("id")

	if id == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "fill query",
		})
	}

	primId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result, err := h.catService.FindCatById(primId)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := helpers.ResponseApi("success retrieve data", result)

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *CatHandler) FindAllCat(c *fiber.Ctx) error {
	limit := c.Query("limit")

	if limit == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "fill query",
		})
	}

	newLimit, err := strconv.Atoi(limit)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result, err := h.catService.FindAllCat(newLimit)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := helpers.ResponseApi("success retrieve data", result)

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *CatHandler) UpdateCatHandler(c *fiber.Ctx) error {
	id := c.Query("id")

	if id == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "fill query",
		})
	}

	primId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, err := h.catService.FindCatById(primId)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	cat := new(models.Cat)

	if err := c.BodyParser(cat); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res.Description = cat.Description
	res.Gender = cat.Gender
	res.Height = cat.Height
	res.Weight = cat.Weight
	res.Medical = cat.Medical
	res.Name = cat.Name
	res.Ras = cat.Ras
	res.UpdatedAt = time.Now()

	result, err := h.catService.UpdateCat(primId, res)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := helpers.ResponseApi("update data success", result+" data updated")

	return c.Status(fiber.StatusOK).JSON(response)

}

func (h *CatHandler) DeleteCatHandler(c *fiber.Ctx) error {
	id := c.Query("id")

	if id == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "fill the query",
		})
	}

	primId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, err := h.catService.DeleteCat(primId)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := helpers.ResponseApi("success delete data", res+" data deleted")

	return c.Status(fiber.StatusOK).JSON(response)
}
