package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	usersTable  = "users"
	drinksTable = "drinks"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgresDB возвращает пул для соединенияс бд
func NewPostgresDB(cfg Config) (*pgxpool.Pool, error) {
	var err error
	dbpool, err := pgxpool.Connect(context.Background(), fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = dbpool.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return dbpool, nil
}
