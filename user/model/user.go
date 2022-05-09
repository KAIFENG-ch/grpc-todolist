package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	services "user/proto/proto"
)

type User struct {
	Username  string `gorm:"not null"`
	Password  string `gorm:"varchar(20);not null"`
	Email     string
	Status    string `gorm:"not null"`
	Birthday  string `gorm:"default:20000101"`
	Signature string `gorm:"default:This man is lazy."`
	gorm.Model
}

const PwdHard = 12

func (receiver User) SetPwd(pwd string) (string, error) {
	pwdByte, err := bcrypt.GenerateFromPassword([]byte(pwd), PwdHard)
	if err != nil {
		return "", err
	}
	receiver.Password = string(pwdByte)
	return receiver.Password, err
}

func (receiver User) CheckPwd(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(receiver.Password), []byte(pwd))
	if err != nil {
		return false
	}
	return true
}

func Build(user User) *services.UserModel {
	userModel := &services.UserModel{
		Id:        uint32(user.ID),
		UserName:  user.Username,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}
	return userModel
}
