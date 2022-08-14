package handler

import (
	"encoding/json"
	routing "github.com/qiangxue/fasthttp-routing"
	"strings"
)

const authHeader = "Authorization"

func (h *Handler) userIdentity(ctx *routing.Context) (err error) {
	encoder := json.NewEncoder(ctx)
	encoder.SetIndent("", "\t")
	header := string(ctx.Request.Header.Peek(authHeader))
	if header == "" {
		err = encoder.Encode(response{
			Description: "empty auth header",
		})
		ctx.SetStatusCode(401)
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		err = encoder.Encode(response{
			Description: "invalid auth header",
		})
		ctx.SetStatusCode(401)
		return
	}
	username, err := h.services.CheckToken(headerParts[1])
	if err != nil {
		err = encoder.Encode(response{
			Description: err.Error(),
		})
		ctx.SetStatusCode(401)
		return
	}
	ctx.Set("username", username)
	return
}
