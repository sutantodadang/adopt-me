package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/sutantodadang/adopt-me/v1/models"
	"github.com/sutantodadang/adopt-me/v1/services"
)

type GalleryHandler struct {
	gallery    services.ServiceGalery
	catService services.ServiceCat
	url        string
	key_image  string
}

func NewGalleryHandler(gallery services.ServiceGalery, catService services.ServiceCat, url, key_image string) *GalleryHandler {
	return &GalleryHandler{gallery, catService, url, key_image}
}

func (h *GalleryHandler) CreateGalleryHandler(c *fiber.Ctx) error {
	id := c.Query("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "fill query",
		})
	}

	primId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	modCat, err := h.catService.FindCatById(primId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := c.Locals("user").(*jwt.Token)
	claim := user.Claims.(jwt.MapClaims)
	user_id := claim["id"].(string)

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	fileHeader := form.File["images"]

	var imageGallery []models.Image
	var gallery models.Gallery

	key := url.QueryEscape(h.key_image)

	fullUrl := fmt.Sprintf("%s%s", h.url, key)

	for _, v := range fileHeader {

		file, err := v.Open()
		if err != nil {
			return err
		}

		buf := bytes.NewBuffer(nil)

		writer := multipart.NewWriter(buf)

		part, err := writer.CreateFormFile("image", v.Filename)

		if err != nil {
			return err
		}

		byt, err := ioutil.ReadAll(file)

		if err != nil {
			return err
		}

		part.Write(byt)
		writer.Close()

		req, err := http.NewRequest("POST", fullUrl, buf)

		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", writer.FormDataContentType())

		client := &http.Client{}

		res, err := client.Do(req)

		if err != nil {
			return err
		}

		var jsonResponse models.ResponseGallery

		err = json.NewDecoder(res.Body).Decode(&jsonResponse)

		if err != nil {
			return err
		}

		catImage := new(models.Image)

		catImage.Id = jsonResponse.Data.Id
		catImage.Filename = jsonResponse.Data.Image.Filename
		catImage.Image_url = jsonResponse.Data.Image.Url
		catImage.Display_url = jsonResponse.Data.Display_url
		catImage.Delete_url = jsonResponse.Data.Delete_url
		catImage.Extension = jsonResponse.Data.Image.Extension
		catImage.Mime = jsonResponse.Data.Image.Mime
		catImage.Thumb = jsonResponse.Data.Thumb.Url

		defer res.Body.Close()

		imageGallery = append(imageGallery, *catImage)
	}

	gallery.Images = imageGallery
	gallery.Cat_Id = modCat.Id.Hex()
	gallery.User_id = user_id

	err = h.gallery.CreateGallery(gallery)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success upload image",
	})
}
