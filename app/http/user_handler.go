package http

import (
	"net/http"
	"rashik/search-scrapper/app/domain"

	"github.com/gin-gonic/gin"
)

type UserHttpHandler struct {
	UserUseCase domain.UserUseCase
}

type UserBasicForm struct {
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
		var json UserBasicForm
		if err := c.ShouldBindJSON(&json); err != nil {

		} else {
			user := &domain.User{
				Email:    &json.Email,
				Password: &json.Password,
			}
			user, err := a.UserUseCase.StoreUser(c, user)
			if err != nil {

			} else {
				c.JSON(http.StatusCreated, user)
			}
		}
	}
}
