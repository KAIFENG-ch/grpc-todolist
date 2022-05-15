package dao

import (
	"user/model"
)

func FindUser(name string) (user *model.User, ok bool) {
	var count int64 = 0
	model.DB.Model(&model.User{}).Where("username=?", name).First(&user).Count(&count)
	if count > 0 {
		return user, true
	}
	return nil, false
}

func CreateUser(user model.User) (err error) {
	err = model.DB.Create(&user).Error
	if err != nil {
		return
	}
	return
}
