package cache

import (
	"errors"
	"gorm.io/gorm"
	"saloon"
	"sync"
)

// Cache - кэш для чтения данных из бд при запуске приложения
type Cache struct {
	data map[string]saloon.User
	mu   sync.RWMutex
}

// NewCache возвращает новую структуру с инициализированным кэшом
func NewCache(db *gorm.DB) (Cache, error) {
	var users []saloon.User
	err := db.Find(&users).Error
	if err != nil {
		return Cache{}, err
	}
	c := Cache{
		data: make(map[string]saloon.User),
		mu:   sync.RWMutex{},
	}
	for _, user := range users {
		c.Put(user)
	}
	return c, nil
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
