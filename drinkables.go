package saloon

type Drink struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Price   string `json:"price"`
	Alcohol int    `json:"alcohol"`
}
