package service

import (
	"github.com/tylorkolbeck/go-cookbook/internal/service/cookbook"
	"github.com/tylorkolbeck/go-cookbook/internal/service/endpoints"
	"github.com/tylorkolbeck/go-cookbook/internal/service/recipe"
	"github.com/tylorkolbeck/go-cookbook/internal/service/user"
)

type AppServices struct {
	UserService      user.UserService
	CookbookService  cookbook.CookbookService
	EndpointsService endpoints.EndpointsService
	RecipeService    recipe.RecipeService
}
