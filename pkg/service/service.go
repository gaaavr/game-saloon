package service

import "saloon/pkg/repository"

// Интерфейс для сущности регистрации и авторизации пользователя
type Authorisation interface {
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
func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
