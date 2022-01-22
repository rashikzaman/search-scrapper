package http

import (
	"net/http"
	"rashik/search-scrapper/app/domain"
	"rashik/search-scrapper/app/logger"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		result, err := a.KeywordUserCase.FetchKeywordsForUser(c, userId, c.Query("search"))
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
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			logger.GetLog().WithFields(logrus.Fields{
				"error": err,
			}).Error("Unable to open the file")
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
