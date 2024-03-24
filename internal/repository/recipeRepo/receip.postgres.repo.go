package recipeRepo

import (
	"github.com/tylorkolbeck/go-cookbook/internal/model"
	"gorm.io/gorm"
)

type PostgresRecipeRepository struct {
	db *gorm.DB
}

func NewPostgresRecipeRepository(db *gorm.DB) *PostgresRecipeRepository {
	return &PostgresRecipeRepository{db: db}
}

func (r *PostgresRecipeRepository) Get() []model.Recipe {
	var recipes []model.Recipe
	r.db.Find(&recipes)

	return recipes
}

func (r *PostgresRecipeRepository) GetByID(id string) (model.Recipe, error) {
	var recipe model.Recipe
	err := r.db.First(&recipe, id).Error

	return recipe, err
}

func (r *PostgresRecipeRepository) Update(recipe_id string, newRecipe model.Recipe, existingRecipe model.Recipe) (model.Recipe, error) {

	err := r.db.Model(&existingRecipe).Updates(newRecipe).Error

	return existingRecipe, err
}

func (r *PostgresRecipeRepository) Delete(recipe_id string) (string, error) {
	err := r.db.Delete(&model.Recipe{}, recipe_id).Error

	return recipe_id, err
}

func (r *PostgresRecipeRepository) Add(recipe model.Recipe) (model.Recipe, error) {
	err := r.db.Create(&recipe).Error

	return recipe, err
}
