package user

import (
	"pathpro-go/dao"
	"pathpro-go/model"
)

func Register(userReq *model.UserRegisterRequest) error {
	db := dao.GetDB()
	user := &dao.User{}
	user.Username = userReq.Username
	user.Password = userReq.Password

	return db.Create(user).Error
}
