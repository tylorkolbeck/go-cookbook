package dto

import (
	"github.com/go-playground/validator/v10"
)

type RecipeBase struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Public       bool   `json:"public"`
	Ingredients  []byte `json:"ingredients"`
	Instructions []byte `json:"instructions"`
}

type CreateRecipeRequest struct {
	RecipeBase
}

type UpdateRecipeRequest struct {
	RecipeBase
}

func (rb *RecipeBase) Validate() error {
	validate := validator.New()
	return validate.Struct(rb)
}

func (r CreateRecipeRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r UpdateRecipeRequest) Validate() error {
	validate := validator.New()
	// Custom validation logic for UpdateRecipeRequest
	return validate.Struct(r)
}
