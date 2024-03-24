package recipeRepo

import (
	"github.com/tylorkolbeck/go-cookbook/internal/model"
)

type RecipeRepository interface {
	Get() []model.Recipe
	GetByID(id string) (model.Recipe, error)
	Update(recipe_id string, newRecipe model.Recipe, existingRecipe model.Recipe) (model.Recipe, error)
	Delete(recipe_id string) (string, error)
	Add(recipe model.Recipe) (model.Recipe, error)
}
