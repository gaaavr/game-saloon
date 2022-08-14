package handler

import (
	"encoding/json"
	routing "github.com/qiangxue/fasthttp-routing"
	"saloon"
)

// Метод для регистрации пользователя
func (h *Handler) register(ctx *routing.Context) (err error) {
	var user saloon.User
	encoder := json.NewEncoder(ctx)
	encoder.SetIndent("", "\t")
	data := ctx.Request.Body()
	err = json.Unmarshal(data, &user)
	if err != nil {
		err = encoder.Encode(failAnswer{Message: err.Error()})
		ctx.Response.SetStatusCode(500)
		return
	}
	id, err := h.services.CreateUser(user)
	if err != nil {
		err = encoder.Encode(failAnswer{Message: err.Error()})
		ctx.Response.SetStatusCode(500)
		return
	}
	err = encoder.Encode(successRegister{id})
	ctx.Response.SetStatusCode(201)
	return
}

// Метод для входа пользователя
func (h *Handler) login(ctx *routing.Context) (err error) {
	var user saloon.User
	encoder := json.NewEncoder(ctx)
	encoder.SetIndent("", "\t")
	data := ctx.Request.Body()
	err = json.Unmarshal(data, &user)
	if err != nil {
		err = encoder.Encode(failAnswer{Message: err.Error()})
		ctx.Response.SetStatusCode(500)
		return
	}
	token, err := h.services.GenerateToken(user.Username, user.Password)
	if err != nil {
		err = encoder.Encode(failAnswer{Message: err.Error()})
		ctx.Response.SetStatusCode(401)
		return
	}
	err = encoder.Encode(successLogin{token})
	ctx.Response.SetStatusCode(200)
	return
}
