package service

import (
	"github.com/tylorkolbeck/go-cookbook/internal/model"
	"github.com/tylorkolbeck/go-cookbook/internal/repository"
)

type EndpointsService struct {
	repo repository.EndpointRepository
}

func NewEndpointsService(repo repository.EndpointRepository) *EndpointsService {
	return &EndpointsService{repo: repo}
}

func (s *EndpointsService) Get() []model.Endpoint {
	return s.repo.Get()
}
