package helper

import (
	"errors"
	"golang-gin3/schema"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secret_token = []byte("mysecretkey")

type JWTClaims struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(user *schema.User) (string, error) {

	claims := JWTClaims{
		user.Id,
		user.Name,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secret_token)

	return ss, err
}

func ValidateToken(header string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(header, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secret_token, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}

		return nil, errors.New("token was expired")
	}

	claims, ok := token.Claims.(*JWTClaims)

	if !ok || !token.Valid {
		return nil, errors.New("token was expired")
	}

	return claims, nil
}
