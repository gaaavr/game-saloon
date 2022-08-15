package repository

import (
	"gorm.io/gorm"
	"saloon"
)

type BarmanPostgres struct {
	db *gorm.DB
}

func NewBarmanPostgres(db *gorm.DB) *BarmanPostgres {
	return &BarmanPostgres{db: db}
}

func (b *BarmanPostgres) CreateDrink(drink saloon.Drink) (id int, err error) {
	err = b.db.Create(&drink).Error
	if err != nil {
		return 0, err
	}
	return drink.Id, nil
}
