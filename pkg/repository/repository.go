package repository

// Интерфейс для сущности регистрации и авторизации пользователя
type Authorisation interface {
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
func NewRepository() *Repository {
	return &Repository{}
}
