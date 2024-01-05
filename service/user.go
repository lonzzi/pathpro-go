package service

import (
	"errors"
	"pathpro-go/dao"
	"pathpro-go/model"
	"pathpro-go/pkg/errno"
	"pathpro-go/utils/jwt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserExist(username string) (bool, error) {
	db := dao.GetDB()
	user := &dao.User{}

	if err := db.Where(&dao.User{Username: username}).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func UserRegister(userReq *model.UserRegisterRequest) error {
	db := dao.GetDB()
	user := &dao.User{}
	user.Username = userReq.Username
	user.Password = userReq.Password

	if exist, err := UserExist(user.Username); err != nil {
		return err
	} else if exist {
		return errno.ErrUserExist
	}

	hashedPassword, err := GeneratePassword(userReq.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	err = db.Create(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errno.ErrUserExist
		}

		return err
	}

	return nil
}

func UserLogin(userReq *model.UserLoginRequest) (*model.UserLoginResponse, error) {
	db := dao.GetDB()
	user := &dao.User{}
	user.Username = userReq.Username
	user.Password = userReq.Password

	if exist, err := UserExist(user.Username); err != nil {
		return nil, err
	} else if !exist {
		return nil, errno.ErrUserNotFound
	}

	err := db.Where("username = ?", userReq.Username).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}
		return nil, err
	}

	if !ValidatePassword(userReq.Password, user.Password) {
		return nil, errno.ErrPasswordIncorrect
	}

	accessToken, refreshToken, err := jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, errno.ErrTokenInvalid
	}

	return &model.UserLoginResponse{
		Id:           user.ID,
		Username:     user.Username,
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func GeneratePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ValidatePassword(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
