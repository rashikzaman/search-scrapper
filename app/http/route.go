package http

import (
	"fmt"
	"rashik/search-scrapper/app/middleware"
	"rashik/search-scrapper/app/repository"
	"rashik/search-scrapper/app/usecase"
	"rashik/search-scrapper/db"

	"github.com/gin-gonic/gin"
)

func InitRouter() {

	r := gin.New()
	r.Use(gin.Logger())
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(500, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(500)
	}))

	r.Static("/public", "./public")

	apiGroup := r.Group("/api")
	{
		userRepository := repository.NewPostgresUserRepository(db.GetDb())
		userUseCase := usecase.NewUserUseCase(userRepository)
		userHandler := NewUserHttpHandler(userUseCase)

		keywordRepository := repository.NewPostgresKeywordRepository(db.GetDb())
		keywordUseCase := usecase.NewKeywordUseCase(keywordRepository)
		keywordHandler := NewKeywordHttpHandler(keywordUseCase)

		authGroup := apiGroup.Group("/auth")
		{
			authGroup.POST("/signup", userHandler.SignUp())
			authGroup.POST("/login", userHandler.Login())
		}

		userGroup := apiGroup.Group("/user", middleware.AuthorizeJwt())
		{
			userGroup.GET("/", userHandler.GetUser())
			userGroup.POST("/keywords", keywordHandler.StoreKeywords())
		}
	}

	r.Run() // listen	 and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
