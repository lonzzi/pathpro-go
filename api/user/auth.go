package user

import (
	"pathpro-go/model"
	"pathpro-go/pkg/engine"
	"pathpro-go/service/user"
)

func Register(ctx *engine.Context) *engine.Response {
	userReq := &model.UserRegisterRequest{}

	err := ctx.Bind(userReq)
	if err != nil {
		return &engine.Response{
			Code: 1,
			Msg:  err.Error(),
		}
	}

	err = user.Register(userReq)
	if err != nil {
		return &engine.Response{
			Code: 1,
			Msg:  err.Error(),
		}
	}

	return &engine.Response{
		Code: 0,
		Msg:  "success",
	}
}

func Login(ctx *engine.Context) *engine.Response {
	return &engine.Response{
		Code: 0,
		Msg:  "pong",
	}
}
