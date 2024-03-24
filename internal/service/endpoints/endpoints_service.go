package endpoints

import (
	"github.com/tylorkolbeck/go-cookbook/internal/model"
	"github.com/tylorkolbeck/go-cookbook/internal/repository/endpointRepo"
)

type EndpointsService struct {
	repo endpointRepo.EndpointRepository
}

func Initialize(repo endpointRepo.EndpointRepository) *EndpointsService {
	return &EndpointsService{repo: repo}
}

func (s *EndpointsService) Get() []model.Endpoint {
	return s.repo.Get()
}
