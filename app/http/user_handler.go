package http

import (
	"net/http"
	"rashik/search-scrapper/app/auth"
	"rashik/search-scrapper/app/domain"
	"rashik/search-scrapper/config"

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
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		} else {
			data := &domain.User{
				Email:    &json.Email,
				Password: &json.Password,
			}
			user, err := a.UserUseCase.StoreUser(c, data)

			if err != nil {
				c.JSON(http.StatusInternalServerError, "Internal Server Error, please try again later")
				return
			} else {
				token, err := auth.CreateJwtAuthToken(int(user.ID), config.GetConfig().GetJwtSecretKey())
				if err != nil {
					c.JSON(http.StatusInternalServerError, "Internal Server Error, please try again later")
					return
				} else {
					c.JSON(http.StatusCreated, gin.H{"access_key": token, "email": user.Email})
					return
				}
			}
		}
	}
}

func (a *UserHttpHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json UserAuthForm
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		} else {
			user, _ := a.UserUseCase.FetchUserByEmail(c, json.Email)
			passwordMatched := a.UserUseCase.CheckPasswordHash(json.Password, *user.Password)
			if passwordMatched {
				token, err := auth.CreateJwtAuthToken(int(user.ID), config.GetConfig().GetJwtSecretKey())
				if err != nil {
					c.JSON(http.StatusInternalServerError, "Internal Server Error, please try again later")
					return
				} else {
					c.JSON(http.StatusCreated, gin.H{"access_key": token})
					return
				}
			}
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
