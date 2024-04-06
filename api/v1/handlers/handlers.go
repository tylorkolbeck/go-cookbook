package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tylorkolbeck/go-cookbook/auth"
	"github.com/tylorkolbeck/go-cookbook/internal/service"
)

type AppHandlers struct {
	CookbookHandler *CookbookHandler
	RecipeHandler   *RecipeHandler
	UserHandler     *UserHandler
}

func RegisterRoutesAndHandlers(router *gin.Engine, services service.AppServices, authConfig *auth.AuthConfig) *AppHandlers {
	appHandlers := &AppHandlers{
		CookbookHandler: RegisterCookbookRoutes(router, services.CookbookService),
		RecipeHandler:   RegisterRecipeRoutes(router, services.RecipeService),
		UserHandler:     RegisterUserRoutes(router, services.UserService, authConfig),
	}

	return appHandlers
}
