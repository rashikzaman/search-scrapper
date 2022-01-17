package middleware

import (
	"net/http"

	"rashik/search-scrapper/app/auth"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizeJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			tokenString := authHeader[len(BEARER_SCHEMA)+1:]
			token, err := auth.ValidateJwtToken(tokenString)
			if err != nil || !token.Valid {
				c.AbortWithStatus(http.StatusUnauthorized)
			} else {
				claims := token.Claims.(jwt.MapClaims)
				c.Set("userId", claims["UserId"])
				c.Next()
			}
		}
	}
}
