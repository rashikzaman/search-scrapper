package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJwtAuthToken(userId int, secretKey string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId":    strconv.Itoa(userId),
		"ExpiresAt": time.Now().Add(time.Hour * 2).Unix(), //expires after two hours
		"IssuedAt":  time.Now().Unix(),
	})
	token, err := claims.SignedString([]byte(secretKey))
	return token, err
}

func ValidateJwtToken(token string, secretKey string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
}
