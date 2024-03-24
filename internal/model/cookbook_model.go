package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CookBook struct {
	gorm.Model
	CookBookID  uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Public      bool      `gorm:"default:false" json:"public"`
}
