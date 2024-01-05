package user

import (
	"pathpro-go/model"
	"pathpro-go/pkg/engine"
	"pathpro-go/pkg/errno"
	"pathpro-go/service"
)

func Register(ctx *engine.Context) *engine.Response {
	userReq := &model.UserRegisterRequest{}

	err := ctx.Bind(userReq)
	if err != nil {
		return &engine.Response{
			Code: errno.ErrBind,
			Msg:  err.Error(),
		}
	}
	err = service.UserRegister(userReq)
	if err != nil {
		errnoErr, ok := err.(errno.ErrCode)
		if ok {
			return &engine.Response{
				Code: errnoErr,
				Msg:  err.Error(),
			}
		} else {
			return &engine.Response{
				Code: errno.InternalServerError,
				Msg:  err.Error(),
			}
		}
	}

	userResp, err := service.UserLogin(&model.UserLoginRequest{
		Username: userReq.Username,
		Password: userReq.Password,
	})
	if err != nil {
		errnoErr, ok := err.(errno.ErrCode)
		if ok {
			return &engine.Response{
				Code: errnoErr,
				Msg:  err.Error(),
			}
		} else {
			return &engine.Response{
				Code: errno.InternalServerError,
				Msg:  err.Error(),
			}
		}
	}

	return &engine.Response{
		Code: errno.OK,
		Msg:  errno.OK.Error(),
		Data: userResp,
	}
}

func Login(ctx *engine.Context) *engine.Response {
	userReq := &model.UserLoginRequest{}

	err := ctx.Bind(userReq)
	if err != nil {
		return &engine.Response{
			Code: errno.ErrBind,
			Msg:  err.Error(),
		}
	}

	userResp, err := service.UserLogin(userReq)
	if err != nil {
		errnoErr, ok := err.(errno.ErrCode)
		if ok {
			return &engine.Response{
				Code: errnoErr,
				Msg:  err.Error(),
			}
		} else {
			return &engine.Response{
				Code: errno.InternalServerError,
				Msg:  err.Error(),
			}
		}
	}

	return &engine.Response{
		Code: errno.OK,
		Msg:  errno.OK.Error(),
		Data: userResp,
	}
}
