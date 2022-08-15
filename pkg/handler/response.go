package handler

import (
	"encoding/json"
	routing "github.com/qiangxue/fasthttp-routing"
)

type response struct {
	Success     bool   `json:"success"`
	Description string `json:"description"`
}

func Response(ctx *routing.Context, statusCode int, description string, status bool) (err error) {
	encoder := json.NewEncoder(ctx)
	encoder.SetIndent("", "\t")
	ctx.Response.SetStatusCode(statusCode)
	return encoder.Encode(response{
		Success:     status,
		Description: description,
	})
}
