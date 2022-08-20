package service

import (
	"errors"
	"math"
	"saloon"
	"saloon/pkg/cache"
	"saloon/pkg/repository"
	"time"
)

type VisitorService struct {
	cache cache.Cache
	repo  repository.Visitor
}

func NewVisitorService(cache cache.Cache, repo repository.Visitor) *VisitorService {
	return &VisitorService{
		repo:  repo,
		cache: cache,
	}
}

// GetData - Метод получения данных посетителя, они берутся из кэша, база не используется
func (v *VisitorService) GetData(username string) (saloon.VisitorData, error) {
	user, err := v.cache.GetUser(username)
	if err != nil {
		return saloon.VisitorData{}, err
	}
	return saloon.VisitorData{
		Dead:  user.Dead,
		Money: user.Money,
		Ppm:   user.Ppm,
	}, nil
}

// BuyDrink - Метод получения данных посетителя, они берутся из кэша, база не используется
func (v *VisitorService) BuyDrink(username, drinkName string) error {
	user, err := v.cache.GetUser(username)
	if err != nil {
		return err
	}
	if user.Dead {
		return errors.New(username + " уже напился до смерти")
	}
	drink, err := v.cache.GetDrink(drinkName)
	if err != nil {
		return err
	}
	if user.Money < drink.Price {
		return errors.New("не хватает денег на покупку")
	}
	ppm := math.Trunc(user.Ppm*100)/100 + math.Trunc(float64(drink.Alcohol))/100
	if user.LastDrink.IsZero() {
		user.Money -= drink.Price
		user.Ppm = math.Round(ppm*100) / 100
		user.LastDrink = time.Now()
	} else {
		diffHours := math.Trunc(time.Since(user.LastDrink).Hours())
		if user.Ppm < diffHours {
			user.Ppm = diffHours
		}
		user.Ppm -= diffHours
		user.Money -= drink.Price
		user.Ppm = math.Round(ppm*100) / 100
		user.LastDrink = time.Now()
	}
	if user.Ppm >= 6 {
		user.Dead = true
	}
	err = v.repo.BuyDrink(user)
	if err != nil {
		return err
	}
	v.cache.PutUser(user)
	return nil
}
