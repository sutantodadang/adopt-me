package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sutantodadang/adopt-me/v1/models"
)

type JwtToken interface {
	GenerateToken(input models.User) (string, error)
}

type jwtService struct {
	secret string
}

func NewJwt(secret string) *jwtService {
	return &jwtService{secret}
}

func (j *jwtService) GenerateToken(input models.User) (string, error) {

	claims := jwt.MapClaims{
		"id":    input.Id,
		"name":  input.Name,
		"place": input.Place,
		"phone": input.Phone,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signToken, err := token.SignedString([]byte(j.secret))

	if err != nil {
		return signToken, err
	}

	return signToken, nil
}
