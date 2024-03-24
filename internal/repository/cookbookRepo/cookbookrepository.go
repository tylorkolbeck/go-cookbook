package cookbookRepo

import (
	"github.com/tylorkolbeck/go-cookbook/internal/model"
)

type CookbookRepository interface {
	Get() []model.CookBook
	GetByID(id string) (model.CookBook, error)
	Update(cookbook_id string, newCookbook model.CookBook, existingCookbook model.CookBook) (model.CookBook, error)
	Delete(cookbook_id string) (string, error)
	Add(cookbook model.CookBook) (model.CookBook, error)
}
