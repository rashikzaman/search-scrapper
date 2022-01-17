package http

import (
	"net/http"
	"rashik/search-scrapper/app/auth"
	"rashik/search-scrapper/app/domain"

	"github.com/gin-gonic/gin"
)

type UserHttpHandler struct {
	UserUseCase domain.UserUseCase
}

type UserAuthForm struct {
	Email    string `form:"email" json:"email" binding:"required,max=255,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

func NewUserHttpHandler(us domain.UserUseCase) *UserHttpHandler {
	handler := &UserHttpHandler{
		UserUseCase: us,
	}
	return handler
}

func (a *UserHttpHandler) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json UserAuthForm
		if err := c.ShouldBindJSON(&json); err != nil {
			token, err := auth.CreateJwtAuthToken()
			if err != nil {
				c.JSON(http.StatusInternalServerError, "Internal Server error, please try again later")
				c.Abort()
			} else {
				c.JSON(http.StatusCreated, gin.H{"access_key": token})
				c.Abort()
			}
		} else {
			c.JSON(http.StatusBadRequest, err)
		}
	}
}

func (a *UserHttpHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json UserAuthForm
		if err := c.ShouldBindJSON(&json); err != nil {

		} else {

		}
	}
}

func (a *UserHttpHandler) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("userId").(string)
		c.JSON(200, gin.H{
			"hello": userId,
		})
	}
}
