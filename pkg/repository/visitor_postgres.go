package repository

import (
	"gorm.io/gorm"
	"saloon"
)

type VisitorPostgres struct {
	db *gorm.DB
}

func NewVisitorPostgres(db *gorm.DB) *VisitorPostgres {
	return &VisitorPostgres{db: db}
}

func (v *VisitorPostgres) BuyDrink(user saloon.User) (err error) {
	err = v.db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
