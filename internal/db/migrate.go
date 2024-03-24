package db

// TODO: Move seed logic out of this file

import (
	"log"

	"github.com/tylorkolbeck/go-cookbook/internal/model"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	db.Exec("DELETE FROM cook_books")
	db.Exec("DELETE FROM users")

	user1 := model.User{
		Email:             "test1@email.com",
		Password:          "password",
		EmailVerified:     true,
		VerificationToken: "123456489",
		Name:              "Test User",
		Role:              "admin",
	}

	if result := db.Create(&user1); result.Error != nil {
		log.Printf("Error seeding user1: %v", result.Error)
		return result.Error
	}

	cookBook1 := model.CookBook{
		Name:        "Chicken Recipes",
		Description: "This cookbook contains all of the chicken recipes I have collected over the years.",
		Public:      true,
		UserId:      user1.ID,
	}

	if result := db.Create(&cookBook1); result.Error != nil {
		log.Printf("Error seeding cookBook1: %v", result.Error)
		return result.Error
	}

	recipe1 := model.Recipe{
		Name:         "Chicken Parmesan",
		Description:  "A classic chicken parmesan recipe.",
		Ingredients:  []byte("Chicken, Parmesan, Tomato Sauce"),
		Instructions: []byte("Bread chicken, fry chicken, add sauce and cheese, bake."),
		Public:       true,
		UserId:       user1.ID,
		CookbookId:   cookBook1.ID,
	}

	if result := db.Create(&recipe1); result.Error != nil {
		log.Printf("Error seeding recipe1: %v", result.Error)
		return result.Error
	}

	return nil
}

func AutoMigrate(db *gorm.DB) error {
	MigrateCookbook(db)
	MigrateUser(db)
	MigrateRecipe(db)

	Seed(db)

	return nil
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{})

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func MigrateCookbook(db *gorm.DB) error {
	err := db.AutoMigrate(&model.CookBook{})

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func MigrateRecipe(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Recipe{})

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}
