package endpointRepo

import (
	"github.com/tylorkolbeck/go-cookbook/api/v1/dto"
	"github.com/tylorkolbeck/go-cookbook/internal/model"
)

type EndpointRepository struct {
	endpoints []model.Endpoint
}

func NewEndpointRepository() *EndpointRepository {
	endpoints := []model.Endpoint{
		// Cookbooks
		{
			Name:   "List Cookbooks",
			Method: "GET",
			Path:   "/cookbooks",
		},
		{
			Name:   "Create Cookbook",
			Method: "POST",
			Path:   "/cookbooks",
			Body:   dto.AddCookbookRequest{},
		},
		{
			Name:   "Get Cookbook",
			Method: "GET",
			Path:   "/cookbooks/:id",
		},
		{
			Name:   "Update Cookbook",
			Method: "PUT",
			Path:   "/cookbooks/:id",
			Body:   dto.UpdateCookbookRequest{},
		},
		{
			Name:   "Delete Cookbook",
			Method: "DELETE",
			Path:   "/cookbooks/:id",
		},

		// Recipes
		{
			Name:   "List Recipes",
			Method: "GET",
			Path:   "/recipes",
		},
		{
			Name:   "Create Recipe",
			Method: "POST",
			Path:   "/recipes",
			Body:   dto.CreateRecipeRequest{},
		},
		{
			Name:   "Get Recipe",
			Method: "GET",
			Path:   "/recipes/:id",
		},
		{
			Name:   "Update Recipe",
			Method: "PUT",
			Path:   "/recipes/:id",
			Body:   dto.UpdateRecipeRequest{},
		},
		{
			Name:   "Delete Recipe",
			Method: "DELETE",
			Path:   "/recipes/:id",
		},

		// Users
		{
			Name:   "List Users",
			Method: "GET",
			Path:   "/users",
		},
		{
			Name:   "Create User",
			Method: "POST",
			Path:   "/users",
			Body:   dto.CreateUserRequest{},
		},
		{
			Name:   "Login",
			Method: "POST",
			Path:   "/login",
			Body:   dto.LoginRequest{},
		},
		{
			Name:   "Verify Email",
			Method: "GET",
			Path:   "/verify/:token",
		},
		{
			Name:   "List Users",
			Method: "GET",
			Path:   "/users",
		},
		{
			Name:   "Get User By ID",
			Method: "GET",
			Path:   "/users/:id",
		},
		{
			Name:   "Delete User",
			Method: "DELETE",
			Path:   "/users/:id",
		},
		{
			Name:   "Update User",
			Method: "PUT",
			Path:   "/users/:id",
			Body:   dto.UpdateUserRequest{},
		},
	}

	return &EndpointRepository{endpoints: endpoints}
}

func (r *EndpointRepository) Get() []model.Endpoint {
	return r.endpoints
}
