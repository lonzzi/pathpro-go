package user

import (
	"pathpro-go/model"
	"pathpro-go/pkg/engine"
	"pathpro-go/pkg/errno"
	"pathpro-go/service"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user with the provided details
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.UserRegisterRequest true "User registration details"
// @Success 200 {object} engine.rawResponse[model.UserLoginResponse] "Successfully registered user"
// @Failure 400 {object} engine.rawResponse[model.UserLoginResponse] "Invalid request"
// @Failure 500 {object} engine.rawResponse[model.UserLoginResponse] "Internal server error"
// @Router /user/register [post]
func Register(ctx *engine.Context) *engine.Response {
	userReq := &model.UserRegisterRequest{}

	err := ctx.Bind(userReq)
	if err != nil {
		return engine.NewErrorResponse(errno.ErrBind)
	}
	err = service.UserRegister(userReq)
	if err != nil {
		errnoErr, ok := err.(errno.ErrCode)
		if ok {
			return engine.NewErrorResponse(errnoErr)
		} else {
			return engine.NewErrorResponse(errno.InternalServerError)
		}
	}

	userResp, err := service.UserLogin(&model.UserLoginRequest{
		Username: userReq.Username,
		Password: userReq.Password,
	})
	if err != nil {
		errnoErr, ok := err.(errno.ErrCode)
		if ok {
			return engine.NewErrorResponse(errnoErr)
		} else {
			return engine.NewErrorResponse(errno.InternalServerError)
		}
	}

	return engine.NewSuccessResponse[*model.UserLoginResponse](userResp)
}

// Login godoc
// @Summary Log in a user
// @Description Log in a user with the provided credentials
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.UserLoginRequest true "User login details"
// @Success 200 {object} engine.rawResponse[model.UserLoginResponse] "Successfully logged in user"
// @Failure 400 {object} engine.rawResponse[model.UserLoginResponse] "Invalid request"
// @Failure 500 {object} engine.rawResponse[model.UserLoginResponse] "Internal server error"
// @Router /user/login [post]
func Login(ctx *engine.Context) *engine.Response {
	userReq := &model.UserLoginRequest{}

	err := ctx.Bind(userReq)
	if err != nil {
		return engine.NewErrorResponse(errno.ErrBind)
	}

	userResp, err := service.UserLogin(userReq)
	if err != nil {
		errnoErr, ok := err.(errno.ErrCode)
		if ok {
			return engine.NewErrorResponse(errnoErr)
		} else {
			return engine.NewErrorResponse(errno.InternalServerError)
		}
	}

	return engine.NewSuccessResponse[*model.UserLoginResponse](userResp)
}
