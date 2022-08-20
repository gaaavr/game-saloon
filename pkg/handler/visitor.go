package handler

import (
	"encoding/json"
	"fmt"
	routing "github.com/qiangxue/fasthttp-routing"
	"saloon"
)

// Метод для получения данных клиента о себе
func (h *Handler) getData(ctx *routing.Context) (err error) {
	if ctx.Response.StatusCode() != 200 {
		return
	}
	username := ctx.Get(keyCtx)
	usernameStr, err := getUsername(username)
	if err != nil {
		err = Response(ctx, 500, err.Error(), false)
		return
	}
	role, err := h.services.CheckRole(usernameStr)
	if err != nil {
		err = Response(ctx, 400, err.Error(), false)
		return
	}
	if role != "visitor" {
		err = Response(ctx, 400, "посмотреть свои данные могут только посетители", false)
		return
	}
	data, err := h.services.GetData(usernameStr)
	if err != nil {
		err = Response(ctx, 500, err.Error(), false)
		return
	}
	encoder := json.NewEncoder(ctx)
	encoder.SetIndent("", "\t")
	ctx.Response.SetStatusCode(200)
	return encoder.Encode(data)
}

// Метод для покупки и выпивания напитка
func (h *Handler) buyDrink(ctx *routing.Context) (err error) {
	if ctx.Response.StatusCode() != 200 {
		return
	}
	username := ctx.Get(keyCtx)
	usernameStr, err := getUsername(username)
	if err != nil {
		err = Response(ctx, 500, err.Error(), false)
		return
	}
	role, err := h.services.CheckRole(usernameStr)
	if err != nil {
		err = Response(ctx, 400, err.Error(), false)
		return
	}
	if role != "visitor" {
		err = Response(ctx, 400, "покупать напитки могут только посетители", false)
		return
	}
	var drink saloon.Drink
	data := ctx.Request.Body()
	err = json.Unmarshal(data, &drink)
	if err != nil {
		err = Response(ctx, 500, err.Error(), false)
		return
	}
	err = h.services.BuyDrink(usernameStr, drink.Name)
	if err != nil {
		err = Response(ctx, 500, err.Error(), false)
		return
	}
	success := fmt.Sprintf("%s выпил напиток %s",
		usernameStr, drink.Name)
	err = Response(ctx, 200, success, true)
	return
}
