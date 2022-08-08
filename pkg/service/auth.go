package service

import (
	"saloon"
	"saloon/pkg/repository"
)

type AuthService struct {
	repo repository.Authorisation
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user saloon.User) (id int, err error) {
	user.Role = "visitor"
	user.Money = 1000
	return a.repo.CreateUser(user)
}
