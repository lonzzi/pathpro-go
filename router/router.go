package router

import (
	"pathpro-go/api/user"
	"pathpro-go/middleware"
	"pathpro-go/pkg/engine"
	"pathpro-go/pkg/errno"
)

func Init(r *engine.Engine) {
	r.GET("/ping", func(ctx *engine.Context) *engine.Response {
		return &engine.Response{
			Code: errno.OK,
			Msg:  "pong",
		}
	})

	userGroup := r.Group("/user")
	{
		userGroup.POST("register", user.Register)
		userGroup.POST("login", user.Login)
	}

	r.Use(middleware.JWTAuth()).GET("/test", func(ctx *engine.Context) *engine.Response {
		return &engine.Response{
			Code: errno.OK,
			Msg:  "test",
		}
	})
}
