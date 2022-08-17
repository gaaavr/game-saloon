package service

import (
	"saloon"
	"saloon/pkg/cache"
	"saloon/pkg/repository"
)

// Интерфейс для сущности регистрации и авторизации пользователя
type Authorisation interface {
	CreateUser(user saloon.User) (string, error)
	GenerateToken(username, password string) (string, error)
	CheckToken(token string) (username string, err error)
}

// Интерфейс для сущности бармена
type Barman interface {
	CreateDrink(drink saloon.Drink) (id int, err error)
	CheckRole(username string) (role string, err error)
	GetDrinks() (list []saloon.Drink)
}

// Интерфейс для сущности посетителя бара
type Visitor interface {
	GetData(username string) (saloon.VisitorData, error)
	BuyDrink(username, drinkName string) error
}

// Собираем все методы, отвечающие за бизнес-логику в одном месте
type Service struct {
	Authorisation
	Barman
	Visitor
}

// Функция конструктор для Service
func NewService(cache cache.Cache, repository *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(cache, repository.Authorisation),
		Barman:        NewBarmanService(cache, repository.Barman),
		Visitor:       NewVisitorService(cache, repository.Visitor),
	}
}
