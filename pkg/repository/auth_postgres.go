package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"saloon"
)

type AuthPostgres struct {
	db *pgxpool.Pool
}

func NewAuthPostgres(db *pgxpool.Pool) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user saloon.User) (id int, err error) {
	query := fmt.Sprintf("INSERT INTO %s (username, password, role, ppm, money, dead, last_drink) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id", usersTable)
	row := a.db.QueryRow(context.Background(), query, user.Username, user.Password, user.Role, user.Ppm, user.Money, user.Dead, user.LastDrink)
	if err = row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
