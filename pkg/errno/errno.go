package errno

import "strconv"

type ErrCode struct {
	Code           int
	HTTPStatusCode int
	Message        string
}

func (e ErrCode) Error() string {
	return e.Message
}

func (e ErrCode) String() string {
	return strconv.Itoa(e.Code)
}

// Server errors
var (
	OK                  = ErrCode{Code: 0, HTTPStatusCode: 200, Message: "ok"}
	InternalServerError = ErrCode{Code: 10001, HTTPStatusCode: 500, Message: "internal server error"}
	ErrBind             = ErrCode{Code: 10002, HTTPStatusCode: 400, Message: "error occurred while binding the request body to the struct"}
	ErrValidation       = ErrCode{Code: 10003, HTTPStatusCode: 400, Message: "validation failed"}
	ErrDatabase         = ErrCode{Code: 10004, HTTPStatusCode: 500, Message: "database error"}
	ErrEncrypt          = ErrCode{Code: 10006, HTTPStatusCode: 500, Message: "encrypt error"}
	ErrDecrypt          = ErrCode{Code: 10007, HTTPStatusCode: 500, Message: "decrypt error"}
)

// JWT errors
var (
	ErrTokenInvalid         = ErrCode{Code: 20000, HTTPStatusCode: 401, Message: "token is invalid"}
	ErrExpiredToken         = ErrCode{Code: 20001, HTTPStatusCode: 401, Message: "token is expired"}
	ErrTokenNotFound        = ErrCode{Code: 20002, HTTPStatusCode: 401, Message: "token not found"}
	ErrRefreshTokenNotFound = ErrCode{Code: 20003, HTTPStatusCode: 401, Message: "refresh token is not found"}
)

// User errors
var (
	ErrUserNotFound      = ErrCode{Code: 21000, HTTPStatusCode: 400, Message: "the user was not found"}
	ErrPasswordIncorrect = ErrCode{Code: 21001, HTTPStatusCode: 400, Message: "password is incorrect"}
	ErrUserExist         = ErrCode{Code: 21002, HTTPStatusCode: 400, Message: "the user was exist"}
)
