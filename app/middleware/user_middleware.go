package middleware

import (
	"net/http"

	"rashik/search-scrapper/app/auth"
	"rashik/search-scrapper/config"

	"github.com/gin-gonic/gin"
)

func AuthorizeJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			tokenString := authHeader[len(BEARER_SCHEMA)+1:]
			token, err := auth.ValidateJwtToken(tokenString, config.GetConfig().GetJwtSecretKey())
			if err != nil || !token.Valid {
				c.JSON(http.StatusUnauthorized, err.Error())
				c.Abort()
			} else {
				if claims, ok := token.Claims.(*auth.CustomClaims); ok && token.Valid {
					c.Set("userId", claims.UserId)
					c.Next()
				} else {
					c.JSON(http.StatusUnauthorized, ok)
					c.Abort()
				}
			}
		}
	}
}
