package cache

import (
	"errors"
	"gorm.io/gorm"
	"saloon"
	"sync"
)

// AuthCache - мнтерфейс, описывающий поведение кэша при авторизации пользователя
type AuthCache interface {
	Put(user saloon.User)
	IsExist(username string) bool
	GetLen() int
	Get(username string) (saloon.User, error)
}

// Cache - кэш для чтения данных из бд при запуске приложения
type Cache struct {
	data map[string]saloon.User
	mu   sync.RWMutex
}

// NewCache возвращает новую структуру с инициализированным кэшом
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]saloon.User),
		mu:   sync.RWMutex{},
	}
}

// GetLen возвращает количество пользователей
func (c *Cache) GetLen() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	l := len(c.data)
	return l
}

// Get возвращает данные пользователя из кэша по его username
func (c *Cache) Get(username string) (saloon.User, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	user, ok := c.data[username]
	if !ok {
		return saloon.User{}, errors.New("the user isn't in cache")
	}

	return user, nil
}

// Put добавляет юзера в кэш
func (c *Cache) Put(user saloon.User) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[user.Username] = user
}

// IsExist проверяет наличие пользователя в кэше по его username
func (c *Cache) IsExist(username string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.data[username]

	return ok
}

// RestoreCache заполняет кэш пользователями из базы данных
func (c *Cache) RestoreCache(db *gorm.DB) error {
	var users []saloon.User
	db.Find(&users)
	for _, user := range users {
		c.Put(user)
	}
	return nil
}
