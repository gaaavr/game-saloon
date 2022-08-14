package service

import (
	"saloon"
	"saloon/pkg/cache"
	"saloon/pkg/repository"
)

// Интерфейс для сущности регистрации и авторизации пользователя
type Authorisation interface {
	CreateUser(user saloon.User) (id int, err error)
	GenerateToken(username, password string) (string, error)
	CheckToken(token string) (username string, err error)
}

// Интерфейс для сущности бармена
type Barman interface {
}

// Интерфейс для сущности посетителя бара
type Visitor interface {
}

// Собираем все методы, отвечающие за бизнес-логику в одном месте
type Service struct {
	Authorisation
	Barman
	Visitor
}

// Функция конструктор для Service
func NewService(cache *cache.Cache, repository *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(cache, repository.Authorisation),
	}
}
