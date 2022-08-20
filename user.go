package saloon

import (
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	Ppm       float64   `json:"ppm"`
	Money     int       `json:"money"`
	Dead      bool      `json:"dead"`
	LastDrink time.Time `json:"last_drink"`
}

type VisitorData struct {
	Dead  bool    `json:"dead"`
	Money int     `json:"money"`
	Ppm   float64 `json:"ppm"`
}
