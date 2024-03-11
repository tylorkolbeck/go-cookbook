package service

import (
	"time"

	"github.com/tylorkolbeck/go-cookbook/api/v1/dto"
	"github.com/tylorkolbeck/go-cookbook/internal/model"
	"github.com/tylorkolbeck/go-cookbook/internal/repository"
)

type RecipeService struct {
	repo repository.RecipeRepository
}

func NewRecipeService(repo repository.RecipeRepository) *RecipeService {
	return &RecipeService{repo: repo}
}

func (s *RecipeService) Get() []model.Recipe {
	return s.repo.Get()
}

func (s *RecipeService) GetByID(id string) (model.Recipe, error) {
	return s.repo.GetByID(id)
}

func (s *RecipeService) Update(recipe_id string, newRecipeValues dto.UpdateRecipeRequest) (model.Recipe, error) {
	existing_recipe, err := s.repo.GetByID(recipe_id)

	if err != nil {
		return model.Recipe{}, err
	}

	updatedRecipe := model.Recipe{
		User_id:      existing_recipe.User_id,
		Name:         newRecipeValues.Name,
		Description:  newRecipeValues.Description,
		Ingredients:  newRecipeValues.Ingredients,
		Instructions: newRecipeValues.Instructions,
		Public:       newRecipeValues.Public,
		Created_at:   existing_recipe.Created_at,
		Updated_at:   time.Now(),
	}

	return s.repo.Update(recipe_id, updatedRecipe, existing_recipe)
}

func (s *RecipeService) Delete(recipe_id string) (string, error) {
	_, err := s.repo.GetByID(recipe_id)

	if err != nil {
		return "", err
	}

	return s.repo.Delete(recipe_id)
}

func (s *RecipeService) Add(recipe dto.CreateRecipeRequest) (model.Recipe, error) {
	newRecipe := model.Recipe{
		User_id:      "1",
		Name:         recipe.Name,
		Description:  recipe.Description,
		Ingredients:  recipe.Ingredients,
		Instructions: recipe.Instructions,
		Public:       recipe.Public,
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
	}

	return s.repo.Add(newRecipe)
}
