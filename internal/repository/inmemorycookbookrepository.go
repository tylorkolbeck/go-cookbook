package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/tylorkolbeck/go-cookbook/internal/model"
)

func NewInMemoryCookbookRepository() *InMemoryCookbookRepository {
	cookbooks := []model.CookBook{
		{Cookbook_id: uuid.New().String(), User_id: "1", Name: "Breakfast", Description: "Breakfast recipes", Public: true, Created_at: time.Now(), Updated_at: time.Now()},
		{Cookbook_id: uuid.New().String(), User_id: "2", Name: "Lunch", Description: "Lunch recipes", Public: true, Created_at: time.Now(), Updated_at: time.Now()},
		{Cookbook_id: uuid.New().String(), User_id: "3", Name: "Dinner", Description: "Dinner recipes", Public: true, Created_at: time.Now(), Updated_at: time.Now()},
		{Cookbook_id: uuid.New().String(), User_id: "3", Name: "Dessert", Description: "Dessert recipes", Public: true, Created_at: time.Now(), Updated_at: time.Now()},
		{Cookbook_id: uuid.New().String(), User_id: "1", Name: "Snacks", Description: "Snack recipes", Public: true, Created_at: time.Now(), Updated_at: time.Now()},
	}

	return &InMemoryCookbookRepository{cookbooks: cookbooks}
}

type InMemoryCookbookRepository struct {
	cookbooks []model.CookBook
}

func (r *InMemoryCookbookRepository) Get() []model.CookBook {
	return r.cookbooks
}

func (r *InMemoryCookbookRepository) Add(cookbook model.CookBook) (model.CookBook, error) {
	cookbook.Cookbook_id = uuid.New().String()
	r.cookbooks = append(r.cookbooks, cookbook)
	return cookbook, nil
}

func (r *InMemoryCookbookRepository) GetByID(id string) (model.CookBook, error) {
	for _, c := range r.cookbooks {
		if c.Cookbook_id == id {
			return c, nil
		}
	}
	return model.CookBook{}, NotFoundError
}

func (r *InMemoryCookbookRepository) Update(cookbook_id string, cookbook model.CookBook, existingCookbook model.CookBook) (model.CookBook, error) {
	for i, c := range r.cookbooks {
		if c.Cookbook_id == cookbook_id {
			// The fields that can be updated for a cookbook are Name, Description, and Public
			r.cookbooks[i].Name = cookbook.Name
			r.cookbooks[i].Description = cookbook.Description
			r.cookbooks[i].Public = cookbook.Public
			return r.cookbooks[i], nil
		}
	}
	return model.CookBook{}, NotFoundError
}

func (r *InMemoryCookbookRepository) Delete(cookbook_id string) (string, error) {
	for i, c := range r.cookbooks {
		if c.Cookbook_id == cookbook_id {
			r.cookbooks = append(r.cookbooks[:i], r.cookbooks[i+1:]...)
			return cookbook_id, nil
		}
	}
	return "", NotFoundError
}
