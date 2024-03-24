package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"ID"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Ingredients  []byte    `json:"ingredients"`
	Instructions []byte    `json:"instructions"`
	Public       bool      `json:"public"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
	UserId       uuid.UUID `json:"user_id"`
	User         User      `json:"user"`
	CookbookId   uuid.UUID `json:"cookbook_id"`
	Cookbook     CookBook  `json:"cookbook"`
}

type RecipeIngredient struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Quantity string    `json:"quantity"`
	Unit     string    `json:"unit"`
	Step     int       `json:"step"`
}

type RecipeInstruction struct {
	ID          uuid.UUID `json:"id"`
	Step        int       `json:"step"`
	Instruction string    `json:"instruction"`
}
