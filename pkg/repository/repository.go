package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"saloon"
)

// Интерфейс для сущности регистрации и авторизации пользователя
type Authorisation interface {
	CreateUser(user saloon.User) (id int, err error)
}

// Интерфейс для сущности бармена
type Barman interface {
}

// Интерфейс для сущности посетителя бара
type Visitor interface {
}

// Собираем все методы, отвечающие за работу с бд в одном месте
type Repository struct {
	Authorisation
	Barman
	Visitor
}

// Функция конструктор для Repository
func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Authorisation: NewAuthPostgres(db),
	}
}
