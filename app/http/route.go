package http

import (
	"fmt"
	"rashik/search-scrapper/app/middleware"
	"rashik/search-scrapper/app/repository"
	"rashik/search-scrapper/app/usecase"
	"rashik/search-scrapper/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
	}))

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
		keywordUseCase := usecase.NewKeywordUseCase(keywordRepository, userUseCase)
		keywordHandler := NewKeywordHttpHandler(keywordUseCase)

		authGroup := apiGroup.Group("/auth")
		{
			authGroup.POST("/signup", userHandler.SignUp())
			authGroup.POST("/login", userHandler.Login())
		}

		userGroup := apiGroup.Group("/user", middleware.AuthorizeJwt())
		{
			userGroup.GET("/", userHandler.GetUser())
			userGroup.GET("/keywords", keywordHandler.FetchUserKeywords())
			userGroup.POST("/keywords", keywordHandler.StoreKeywords())
		}
	}

	return r
}
