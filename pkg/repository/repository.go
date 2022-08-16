package repository

import (
	"gorm.io/gorm"
	"saloon"
)

// Интерфейс для сущности регистрации и авторизации пользователя
type Authorisation interface {
	CreateUser(user saloon.User) (id int, err error)
	GetUser(username, password string) (saloon.User, error)
}

// Интерфейс для сущности бармена
type Barman interface {
	CreateDrink(drink saloon.Drink) (id int, err error)
}

// Интерфейс для сущности посетителя бара
type Visitor interface {
	BuyDrink(user saloon.User) (err error)
}

// Собираем все методы, отвечающие за работу с бд в одном месте
type Repository struct {
	Authorisation
	Barman
	Visitor
}

// Функция конструктор для Repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorisation: NewAuthPostgres(db),
		Barman:        NewBarmanPostgres(db),
		Visitor:       NewVisitorPostgres(db),
	}
}
