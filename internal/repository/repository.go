package repository

import "gorm.io/gorm"

type AppRepository struct {
	CookbookRepository CookbookRepository
	RecipeRepository   RecipeRepository
	UserRepository     UserRepository
}

func NewAppRepository(db *gorm.DB) (AppRepository, error) {
	appRepository := AppRepository{
		CookbookRepository: NewPostgresCookbookRepository(db),
		RecipeRepository:   NewPostgresRecipeRepository(db),
		UserRepository:     NewPostgresUserRepository(db),
	}

	// dbErr := db.AutoMigrate(db)
	// if dbErr != nil {
	// 	return AppRepository{}, dbErr
	// }

	return appRepository, nil
}
