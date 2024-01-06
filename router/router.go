package router

import (
	"pathpro-go/api/user"
	"pathpro-go/middleware"
	"pathpro-go/pkg/engine"
	"pathpro-go/pkg/errno"

	swaggerFiles "github.com/swaggo/files" // swagger embed files
)

func Init(e *engine.Engine) {
	e.GET("/ping", func(ctx *engine.Context) *engine.Response {
		return engine.NewResponseWithMsg(errno.OK, "pong", nil)
	})

	e.GET("/swagger/*any", e.WrapHandler(swaggerFiles.Handler)) // swagger embed router

	userGroup := e.Group("/user")
	{
		userGroup.POST("register", user.Register)
		userGroup.POST("login", user.Login)
	}

	e.Use(middleware.JWTAuth()).GET("/test", func(ctx *engine.Context) *engine.Response {
		return engine.NewResponseWithMsg(errno.OK, "test", nil)
	})
}
