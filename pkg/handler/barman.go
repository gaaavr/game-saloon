package handler

import (
	"encoding/json"
	routing "github.com/qiangxue/fasthttp-routing"
	"saloon"
)

// Метод для создания напитка барменом
func (h *Handler) createDrink(ctx *routing.Context) (err error) {
	var drink saloon.Drink
	encoder := json.NewEncoder(ctx)
	encoder.SetIndent("", "\t")
	data := ctx.Request.Body()
	err = json.Unmarshal(data, &drink)
	if err != nil {
		err = encoder.Encode(response{
			Description: err.Error(),
		})
		ctx.Response.SetStatusCode(500)
		return
	}
	if drink.Name == "" {
		err = encoder.Encode(response{
			Description: "у напитка должно быть название",
		})
		ctx.Response.SetStatusCode(400)
		return
	}
	if drink.Alcohol < 0 || drink.Alcohol > 100 {
		err = encoder.Encode(response{
			Description: "алкогольность должна быть от 0 до 100%",
		})
		ctx.Response.SetStatusCode(400)
		return
	}
	return
}
