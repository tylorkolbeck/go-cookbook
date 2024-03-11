package model

import "time"

type Recipe struct {
	Recipe_id    string              `json:"recipe_id"`
	User_id      string              `json:"user_id"`
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	Ingredients  []RecipeIngredient  `json:"ingredients"`
	Instructions []RecipeInstruction `json:"instructions"`
	Public       bool                `json:"public"`
	Created_at   time.Time           `json:"created_at"`
	Updated_at   time.Time           `json:"updated_at"`
}

type RecipeIngredient struct {
	Ingredient_id string `json:"ingredient_id"`
	Name          string `json:"name"`
	Quantity      string `json:"quantity"`
	Unit          string `json:"unit"`
}

type RecipeInstruction struct {
	Instruction_id string `json:"instruction_id"`
	Step           int    `json:"step"`
	Instruction    string `json:"instruction"`
}
