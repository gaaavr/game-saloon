package repository

import (
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
	a.db.Create(&user)
	return user.Id, nil
}
