package db

import (
	"github.com/tylorkolbeck/go-cookbook/internal/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	MigrateCookbook(db)

	return nil
}

func MigrateCookbook(db *gorm.DB) error {
	err := db.AutoMigrate(&model.CookBook{})

	if err != nil {
		return err
	}

	// Create
	db.Create(&model.CookBook{
		Name:        "Chicken Recipes",
		Description: "This cookbook contains all of the chicken recipes I have collected over the years.",
		Public:      true,
	})

	if err != nil {
		return err
	}

	return nil
}
