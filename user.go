package saloon

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Ppm      int    `json:"ppm"`
	Money    int    `json:"money"`
	Status   string `json:"status"`
}
