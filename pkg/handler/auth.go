package handler

import (
	"encoding/json"
	"fmt"
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
		err = Response(ctx, 500, err.Error(), false)
		return
	}
	if user.Username == "" || user.Password == "" {
		err = Response(ctx, 400, "username и password не должны быть пустыми", false)
		return
	}
	id, err := h.services.CreateUser(user)
	if err != nil {
		err = Response(ctx, 500, err.Error(), false)
		return
	}
	success := fmt.Sprintf("Пользователь %s успешно зарегистрирован, id=%d, role=%s", user.Username, id, user.Role)
	err = Response(ctx, 201, success, true)
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
		err = Response(ctx, 500, err.Error(), false)
		return
	}
	token, err := h.services.GenerateToken(user.Username, user.Password)
	if err != nil {
		err = Response(ctx, 401, err.Error(), false)
		return
	}
	err = Response(ctx, 200, token, true)
	return
}
