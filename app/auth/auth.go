package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func CreateJwtAuthToken(userId int, secretKey string) (string, error) {
	claims := CustomClaims{
		strconv.Itoa(userId),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "scrapper",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secretKey))
	return ss, err
}

func ValidateJwtToken(token string, secretKey string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})
}
