package http

import (
	"fmt"
	"net/http"
	"rashik/search-scrapper/app/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KeywordHttpHandler struct {
	KeywordUserCase domain.KeywordUseCase
}

func NewKeywordHttpHandler(kw domain.KeywordUseCase) *KeywordHttpHandler {
	handler := &KeywordHttpHandler{
		KeywordUserCase: kw,
	}
	return handler
}

func (a *KeywordHttpHandler) FetchUserKeywords() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.MustGet("userId").(string)
		userId, err := strconv.Atoi(id)
		result, err := a.KeywordUserCase.FetchKeywordsForUser(c, userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}

func (a *KeywordHttpHandler) StoreKeywords() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.MustGet("userId").(string)
		userId, err := strconv.Atoi(id)
		file, _, err := c.Request.FormFile("Filename")
		if err != nil {
			fmt.Println("Unable to open the file", err)
		}
		result, err := a.KeywordUserCase.StoreKeywordsFromFile(c, file, userId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusCreated, result)
		}

	}
}
