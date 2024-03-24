package cookbook

import (
	"github.com/tylorkolbeck/go-cookbook/api/v1/dto"
	"github.com/tylorkolbeck/go-cookbook/internal/model"
	"github.com/tylorkolbeck/go-cookbook/internal/repository/cookbookRepo"
)

type CookbookService struct {
	repo cookbookRepo.CookbookRepository
}

func Initialize(repo cookbookRepo.CookbookRepository) *CookbookService {
	return &CookbookService{repo: repo}
}

func (s *CookbookService) Get() []model.CookBook {
	return s.repo.Get()
}

func (s *CookbookService) GetByID(id string) (model.CookBook, error) {
	return s.repo.GetByID(id)
}

func (s *CookbookService) Add(cookbook dto.AddCookbookRequest) (model.CookBook, error) {
	newCookbook := model.CookBook{
		Name:        cookbook.Name,
		Description: cookbook.Description,
		Public:      cookbook.Public,
	}
	return s.repo.Add(newCookbook)
}

func (s *CookbookService) Update(cookbook_id string, newCookbookValues model.CookBook) (model.CookBook, error) {
	existing_cookbook, err := s.repo.GetByID(cookbook_id)

	if err != nil {
		return model.CookBook{}, err
	}

	return s.repo.Update(cookbook_id, newCookbookValues, existing_cookbook)
}

func (s *CookbookService) Delete(cookbook_id string) (string, error) {
	_, err := s.repo.GetByID(cookbook_id)

	if err != nil {
		return "", err
	}

	return s.repo.Delete(cookbook_id)
}
