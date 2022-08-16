package service

import (
	"errors"
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
	if b.cache.DrinkIsExist(drink.Name) {
		return 0, errors.New("такой напиток уже существует")
	}
	id, err = b.repo.CreateDrink(drink)
	if err != nil {
		return 0, err
	}
	drink.Id = id
	b.cache.PutDrink(drink)
	return id, err

}

func (b *BarmanService) CheckRole(username string) (role string, err error) {
	if !b.cache.UserIsExist(username) {
		return "", errors.New("пользователь не найден")
	}
	user, err := b.cache.GetUser(username)
	if err != nil {
		return "", err
	}
	if user.Role == "" {
		return "", errors.New("у пользователя нет роли")
	}
	return user.Role, nil
}
