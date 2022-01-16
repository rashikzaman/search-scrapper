package http

import (
	"net/http"
	"rashik/search-scrapper/app/domain"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUseCase domain.UserUseCase
}

type UserBasicForm struct {
	Email string `form:"email" json:"email" binding:"required,max=255,email"`
}

func NewUserHandler(us domain.UserUseCase) *UserHandler {
	handler := &UserHandler{
		UserUseCase: us,
	}
	return handler
}

func (a *UserHandler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json UserBasicForm
		if err := c.ShouldBindJSON(&json); err != nil {

		} else {
			user := &domain.User{
				Email: &json.Email,
			}
			user, err := a.UserUseCase.StoreUser(c, user)
			if err != nil {

			} else {
				c.JSON(http.StatusCreated, user)
			}
		}
	}
}
