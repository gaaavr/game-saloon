package cache

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"saloon"
	"sync"
)

// Cache - кэш для чтения данных из бд при запуске приложения
type Cache struct {
	data map[int]saloon.User
	mu   sync.RWMutex
}

// NewCache возвращает новую структуру с инициализированным кэшом
func NewCache() *Cache {
	return &Cache{
		data: make(map[int]saloon.User),
		mu:   sync.RWMutex{},
	}
}

// Get возвращает данные пользователя из кэша по его id
func (c *Cache) Get(id int) (saloon.User, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	user, ok := c.data[id]
	if !ok {
		return saloon.User{}, errors.New("the user isn't in cache")
	}

	return user, nil
}

// Put добавляет юзера в кэш
func (c *Cache) Put(user saloon.User) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[user.Id] = user
}

// IsExist проверяет наличие пользователя в кэше по его id
func (c *Cache) IsExist(id int) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, ok := c.data[id]

	return ok
}

// RestoreCache заполняет кэш пользователями из базы данных
func (c *Cache) RestoreCache(db *pgxpool.Pool) error {
	rows, err := db.Query(context.Background(), "select * from users")
	if err != nil {
		return err
	}

	for rows.Next() {
		var user saloon.User
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.Ppm, &user.Money, &user.Dead, &user.LastDrink)
		if err != nil {
			return err
		}

		c.Put(user)
	}

	return rows.Err()
}
