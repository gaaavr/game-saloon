package service

import (
	"saloon"
	"saloon/pkg/cache"
	"saloon/pkg/repository"
)

type BarmanService struct {
	cache cache.Cache
	repo  repository.Barman
}

func NewBarmanService(cache cache.Cache, repo repository.Barman) *BarmanService {
	return &BarmanService{
		repo:  repo,
		cache: cache,
	}
}

func (b *BarmanService) CreateDrink(drink saloon.Drink) (id int, err error) {
	return b.repo.CreateDrink(drink)

}
