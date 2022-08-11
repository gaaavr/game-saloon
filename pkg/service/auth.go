package service

import (
	"fmt"
	"saloon"
	"saloon/pkg/cache"
	"saloon/pkg/repository"
)

type AuthService struct {
	cache cache.AuthCache
	repo  repository.Authorisation
}

func NewAuthService(cache cache.AuthCache, repo repository.Authorisation) *AuthService {
	return &AuthService{
		repo:  repo,
		cache: cache,
	}
}

func (a *AuthService) CreateUser(user saloon.User) (id int, err error) {
	if a.cache.IsExist(user.Username) {
		return 0, fmt.Errorf("такой username уже занят")
	}
	if a.cache.GetLen() == 0 {
		user.Role = "barman"
	} else {
		user.Role = "visitor"
	}
	user.Money = 1000
	id, err = a.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}
	user.Id = id
	a.cache.Put(user)
	return id, err
}
