package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"saloon"
	"saloon/pkg/cache"
	"saloon/pkg/repository"
	"time"
)

const (
	tokenExpire = 1 * time.Hour
	secretKey   = "secret"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	Username string `json:"user_name"`
}
type AuthService struct {
	cache cache.Cache
	repo  repository.Authorisation
}

func NewAuthService(cache cache.Cache, repo repository.Authorisation) *AuthService {
	return &AuthService{
		repo:  repo,
		cache: cache,
	}
}

func (a *AuthService) CreateUser(user saloon.User) (id int, err error) {
	if a.cache.IsExist(user.Username) {
		return 0, fmt.Errorf("такой username уже занят")
	}
	if a.cache.GetLen() == 0 {
		user.Role = "barman"
	} else {
		user.Role = "visitor"
	}
	user.Money = 1000
	id, err = a.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}
	user.Id = id
	a.cache.Put(user)
	return id, err
}

func (a *AuthService) GenerateToken(username, password string) (t string, err error) {
	var user saloon.User
	if !a.cache.IsExist(username) {
		user, err = a.repo.GetUser(username, password)
		if err != nil {
			return "", err
		}
	} else {
		user, err = a.cache.Get(username)
		if err != nil {
			return "", err
		}
		if user.Password != password {
			return "", fmt.Errorf("пользователь с таким именем и паролем не найден")
		}
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Username: user.Username,
	})

	return token.SignedString([]byte(secretKey))
}

func (a *AuthService) CheckToken(token string) (username string, err error) {
	parseToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		err = errors.New("error parsing the token")
		return
	}
	claims, ok := parseToken.Claims.(*tokenClaims)
	if !ok {
		err = errors.New("invalid claims")
		return
	}
	username = claims.Username
	return
}
