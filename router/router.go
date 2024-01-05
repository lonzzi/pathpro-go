package router

import (
	"pathpro-go/api/user"
	"pathpro-go/pkg/engine"
)

func Init(r *engine.Engine) {
	r.GET("/ping", func(ctx *engine.Context) *engine.Response {
		return &engine.Response{
			Code: 0,
			Msg:  "pong",
		}
	})

	r.POST(("user/register"), user.Register)
}
