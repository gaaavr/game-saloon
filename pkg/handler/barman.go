package handler

import (
	"encoding/json"
	"fmt"
	routing "github.com/qiangxue/fasthttp-routing"
	"saloon"
)

// Метод для создания напитка барменом
func (h *Handler) createDrink(ctx *routing.Context) (err error) {
	if ctx.Response.StatusCode() != 200 {
		return
	}
	username := ctx.Get(keyCtx)
	usernameStr, err := getUsername(username)
	if err != nil {
		err = Response(ctx, 500, err.Error(), false)
		return
	}
	var drink saloon.Drink
	encoder := json.NewEncoder(ctx)
	encoder.SetIndent("", "\t")
	data := ctx.Request.Body()
	err = json.Unmarshal(data, &drink)
	if err != nil {
		err = Response(ctx, 500, err.Error(), false)
		return
	}
	if drink.Name == "" {
		err = Response(ctx, 400, "у напитка должно быть название", false)
		return
	}
	if drink.Alcohol < 0 || drink.Alcohol > 100 {
		err = Response(ctx, 400, "алкогольность должна быть от 0 до 100%", false)
		return
	}
	if drink.Price < 0 {
		err = Response(ctx, 400, "цена не может быть отрицательной", false)
		return
	}
	id, err := h.services.CreateDrink(drink)
	if err != nil {
		err = Response(ctx, 500, err.Error(), false)
		return
	}
	success := fmt.Sprintf("%s успешно создал напиток %s с id %d, алкогольностью %d%%,  ценой %d",
		usernameStr, drink.Name, id, drink.Alcohol, drink.Price)
	err = Response(ctx, 201, success, true)
	return
}
