package http

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tylorkolbeck/go-cookbook/api/v1/handlers/cookbookHandler"
	"github.com/tylorkolbeck/go-cookbook/internal/repository/cookbookRepo"
	"github.com/tylorkolbeck/go-cookbook/internal/service/cookbook"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:4200"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.Static("/static", "./public/playground")

	// COOKBOOKS
	cookbookRepo := cookbookRepo.NewInmemoryCookbookRepository()
	cookbookService := cookbook.Initialize(cookbookRepo)
	cookbookHandler := cookbookHandler.NewCookbookHandler(cookbookService)
	router.GET("/cookbooks", cookbookHandler.ListCookbooks)
	router.POST("/cookbooks", cookbookHandler.CreateCookbook)
	router.GET("/cookbooks/:id", cookbookHandler.GetCookbook)
	router.PUT("/cookbooks/:id", cookbookHandler.UpdateCookbook)
	router.DELETE("/cookbooks/:id", cookbookHandler.DeleteCookbook)

	return router
}
