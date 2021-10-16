package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

type MiddlewareService interface {
	AuthMiddle() func(*fiber.Ctx) error
}

type middlewareService struct {
	secret string
}

func NewAuthMiddleware(secret string) *middlewareService {
	return &middlewareService{secret}
}

func (m *middlewareService) AuthMiddle() func(*fiber.Ctx) error {

	return jwtware.New(jwtware.Config{
		// SuccessHandler: func(c *fiber.Ctx) error {
		// 	return c.Next()
		// },
		SigningKey: []byte(m.secret),
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": e.Error()})

		},
		AuthScheme:  "Bearer",
		TokenLookup: "header:Authorization",
	})
}
