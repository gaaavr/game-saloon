package repository

import (
	"fmt"
	"gorm.io/gorm"
	"saloon"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user saloon.User) (id int, err error) {
	err = a.db.Create(&user).Error
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (a *AuthPostgres) GetUser(username, password string) (saloon.User, error) {
	var user saloon.User
	if res := a.db.Where("username = ? AND password = ?", username, password).Find(&user).RowsAffected; res == 0 {
		return user, fmt.Errorf("пользователь с таким именем и паролем не найден")
	}
	return user, nil
}
