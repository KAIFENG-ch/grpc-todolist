package dao

import (
	"user/model"
)

func FindUser(name string) (user model.User, err error) {
	var count int64 = 0
	if err := model.DB.Model(&model.User{}).Where("username=?", name).Count(&count).Error; err != nil {
		return
	}
	user = model.User{
		Username: name,
	}
	return
}

func CreateUser(user model.User) (err error) {
	err = model.DB.Create(&user).Error
	if err != nil {
		return
	}
	return
}
