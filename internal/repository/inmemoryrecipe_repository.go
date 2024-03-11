package repository

import (
	"github.com/google/uuid"
	"github.com/tylorkolbeck/go-cookbook/internal/model"
)

func NewInMemoryRecipeRepository() *InMemoryRecipeRepository {
	return &InMemoryRecipeRepository{
		recipes: []model.Recipe{
			{
				Recipe_id:   uuid.New().String(),
				User_id:     uuid.New().String(),
				Name:        "Pasta",
				Description: "Pasta with red sauce",
				Ingredients: []model.RecipeIngredient{
					{
						Ingredient_id: uuid.New().String(),
						Name:          "Pasta",
						Quantity:      "1",
						Unit:          "lb",
					},
				},
				Instructions: []model.RecipeInstruction{
					{
						Instruction_id: uuid.New().String(),
						Step:           1,
						Instruction:    "Boil water",
					},
				},
				Public: true,
			},
			{
				Recipe_id:   uuid.New().String(),
				User_id:     uuid.New().String(),
				Name:        "Pizza",
				Description: "Pizza with pepperoni",
				Ingredients: []model.RecipeIngredient{
					{
						Ingredient_id: uuid.New().String(),
						Name:          "Pizza dough",
						Quantity:      "1",
						Unit:          "lb",
					},
				},
				Instructions: []model.RecipeInstruction{
					{
						Instruction_id: uuid.New().String(),
						Step:           1,
						Instruction:    "Roll out dough",
					},
				},
				Public: true,
			},
			{
				Recipe_id:   uuid.New().String(),
				User_id:     uuid.New().String(),
				Name:        "Salad",
				Description: "Salad with ranch dressing",
				Ingredients: []model.RecipeIngredient{
					{
						Ingredient_id: uuid.New().String(),
						Name:          "Lettuce",
						Quantity:      "1",
						Unit:          "head",
					},
				},
				Instructions: []model.RecipeInstruction{
					{
						Instruction_id: uuid.New().String(),
						Step:           1,
						Instruction:    "Wash lettuce",
					},
				},
				Public: true,
			},
		},
	}
}

type InMemoryRecipeRepository struct {
	recipes []model.Recipe
}

func (r *InMemoryRecipeRepository) Get() []model.Recipe {
	return r.recipes
}

func (r *InMemoryRecipeRepository) Add(recipe model.Recipe) (model.Recipe, error) {
	recipe.Recipe_id = uuid.New().String()
	r.recipes = append(r.recipes, recipe)
	return recipe, nil
}

func (r *InMemoryRecipeRepository) GetByID(id string) (model.Recipe, error) {
	for _, c := range r.recipes {
		if c.Recipe_id == id {
			return c, nil
		}
	}
	return model.Recipe{}, NotFoundError
}

func (r *InMemoryRecipeRepository) Update(recipe_id string, recipe model.Recipe, existingRecipe model.Recipe) (model.Recipe, error) {
	for i, c := range r.recipes {
		if c.Recipe_id == recipe_id {
			// The fields that can be updated for a recipe are Name, Description, Ingredients, Instructions, and Public
			r.recipes[i].Name = recipe.Name
			r.recipes[i].Description = recipe.Description
			r.recipes[i].Ingredients = recipe.Ingredients
			r.recipes[i].Instructions = recipe.Instructions
			r.recipes[i].Public = recipe.Public
			return r.recipes[i], nil
		}
	}
	return model.Recipe{}, NotFoundError
}

func (r *InMemoryRecipeRepository) Delete(recipe_id string) (string, error) {
	for i, c := range r.recipes {
		if c.Recipe_id == recipe_id {
			r.recipes = append(r.recipes[:i], r.recipes[i+1:]...)
			return recipe_id, nil
		}
	}
	return "", NotFoundError
}
