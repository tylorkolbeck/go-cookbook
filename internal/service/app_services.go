package service

import (
	"github.com/tylorkolbeck/go-cookbook/auth"
	"github.com/tylorkolbeck/go-cookbook/internal/repository"
)

type AppServices struct {
	UserService     *UserService
	CookbookService *CookbookService
	RecipeService   *RecipeService
}

func NewAppServices(repo repository.AppRepository, authConfig auth.AuthConfig) *AppServices {
	return &AppServices{
		UserService:     NewUserService(repo.UserRepository, authConfig),
		CookbookService: NewCookbookService(repo.CookbookRepository),
		RecipeService:   NewRecipeService(repo.RecipeRepository),
	}
}
