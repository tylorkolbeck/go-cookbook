package cookbookRepo

import (
	"github.com/tylorkolbeck/go-cookbook/internal/model"
	"gorm.io/gorm"
)

type PostgresCookbookRepository struct {
	db *gorm.DB
}

func NewPostgresCookbookRepository(db *gorm.DB) *PostgresCookbookRepository {
	return &PostgresCookbookRepository{db: db}
}

func (r *PostgresCookbookRepository) Get() []model.CookBook {
	var cookbooks []model.CookBook
	r.db.Find(&cookbooks)

	return cookbooks
}

func (r *PostgresCookbookRepository) GetByID(id string) (model.CookBook, error) {
	var cookbook model.CookBook
	err := r.db.First(&cookbook, id).Error

	return cookbook, err
}

func (r *PostgresCookbookRepository) Update(cookbook_id string, newCookbook model.CookBook, existingCookbook model.CookBook) (model.CookBook, error) {
	err := r.db.Model(&existingCookbook).Updates(newCookbook).Error

	return existingCookbook, err
}

func (r *PostgresCookbookRepository) Delete(cookbook_id string) (string, error) {
	err := r.db.Delete(&model.CookBook{}, cookbook_id).Error

	return cookbook_id, err
}

func (r *PostgresCookbookRepository) Add(cookbook model.CookBook) (model.CookBook, error) {
	err := r.db.Create(&cookbook).Error

	return cookbook, err
}
