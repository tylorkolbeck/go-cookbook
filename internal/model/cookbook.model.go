package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CookBook struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"ID"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Public      bool      `gorm:"default:false" json:"public"`
	UserId      uuid.UUID `gorm:"type:uuid;not null" json:"userId"`
	User        User      `json:"user"`
}
