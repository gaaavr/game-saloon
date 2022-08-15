package handler

import (
	"encoding/json"
	"errors"
	routing "github.com/qiangxue/fasthttp-routing"
	"strings"
)

const (
	authHeader = "Authorization"
	keyCtx     = "username"
)

func (h *Handler) userIdentity(ctx *routing.Context) (err error) {
	encoder := json.NewEncoder(ctx)
	encoder.SetIndent("", "\t")
	header := string(ctx.Request.Header.Peek(authHeader))
	if header == "" {
		err = Response(ctx, 401, "empty auth header", false)
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		err = Response(ctx, 401, "invalid auth header", false)
		return
	}
	username, err := h.services.CheckToken(headerParts[1])
	if err != nil {
		err = Response(ctx, 401, err.Error(), false)
		return
	}
	ctx.Set("username", username)
	return
}

func getUsername(username interface{}) (string, error) {
	if username == nil {
		err := errors.New("username не найден")
		return "", err
	}
	usernameStr, ok := username.(string)
	if !ok {
		err := errors.New("username задан неправильным типом")
		return "", err
	}
	return usernameStr, nil
}
