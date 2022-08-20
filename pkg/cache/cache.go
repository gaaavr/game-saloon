package cache

import (
	"errors"
	"gorm.io/gorm"
	"saloon"
	"sync"
)

// Cache - кэш для чтения данных из бд при запуске приложения
type Cache struct {
	users  map[string]saloon.User
	drinks map[string]saloon.Drink
	mu     sync.RWMutex
}

// NewCache возвращает новую структуру с инициализированным кэшом
func NewCache(db *gorm.DB) (Cache, error) {
	var users []saloon.User
	err := db.Find(&users).Error
	if err != nil {
		return Cache{}, err
	}
	var drinks []saloon.Drink
	err = db.Find(&drinks).Error
	if err != nil {
		return Cache{}, err
	}
	c := Cache{
		users:  make(map[string]saloon.User),
		drinks: make(map[string]saloon.Drink),
		mu:     sync.RWMutex{},
	}
	for _, user := range users {
		c.PutUser(user)
	}
	for _, drink := range drinks {
		c.PutDrink(drink)
	}
	return c, nil
}

// GetLen возвращает количество пользователей
func (c *Cache) GetUsersLen() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	l := len(c.users)
	return l
}

// GetUser возвращает данные пользователя из кэша по его username
func (c *Cache) GetUser(username string) (saloon.User, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	user, ok := c.users[username]
	if !ok {
		return saloon.User{}, errors.New("the user isn't in cache")
	}

	return user, nil
}

// GetDrink возвращает данные напитка из кэша по его name
func (c *Cache) GetDrink(name string) (saloon.Drink, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	drink, ok := c.drinks[name]
	if !ok {
		return saloon.Drink{}, errors.New("the drink isn't in cache")
	}

	return drink, nil
}

// PutUser добавляет юзера в кэш
func (c *Cache) PutUser(user saloon.User) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.users[user.Username] = user
}

// PutDrink добавляет напиток в кэш
func (c *Cache) PutDrink(drink saloon.Drink) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.drinks[drink.Name] = drink
}

// UserIsExist проверяет наличие пользователя в кэше по его username
func (c *Cache) UserIsExist(username string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.users[username]

	return ok
}

// DrinkIsExist проверяет наличие напитка в кэше по его name
func (c *Cache) DrinkIsExist(name string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.drinks[name]

	return ok
}

// ListDrinks проверяет наличие напитка в кэше по его name
func (c *Cache) ListDrinks() []saloon.Drink {
	list := make([]saloon.Drink, 0, len(c.drinks))
	c.mu.RLock()
	defer c.mu.RUnlock()
	for _, drink := range c.drinks {
		list = append(list, drink)
	}

	return list
}
