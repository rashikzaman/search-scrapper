package http

import (
	"net/http"
	"rashik/search-scrapper/app/auth"
	"rashik/search-scrapper/app/domain"
	"rashik/search-scrapper/config"
	"strconv"

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
			user, err := a.UserUseCase.FetchUserByEmail(c, json.Email)
			if err != nil {
				c.JSON(http.StatusUnauthorized, "Email or Password Doesn't match")
			}
			passwordMatched := a.UserUseCase.CheckPasswordHash(json.Password, *user.Password)
			if passwordMatched {
				token, err := auth.CreateJwtAuthToken(int(user.ID), config.GetConfig().GetJwtSecretKey())
				if err != nil {
					c.JSON(http.StatusInternalServerError, "Internal Server Error, please try again later")
					return
				} else {
					c.JSON(http.StatusOK, gin.H{"access_key": token, "email": user.Email})
					return
				}
			} else {
				c.JSON(http.StatusUnauthorized, "Email or Password Doesn't match")
			}
		}
	}
}

func (a *UserHttpHandler) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.MustGet("userId").(string)
		userId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Internal Server Error, please try again later")
			return
		}
		user, err := a.UserUseCase.FetchUserById(c, userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Internal Server Error, please try again later")
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
