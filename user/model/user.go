package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	services "user/proto/pb"
)

type User struct {
	Username  string `gorm:"not null"`
	Password  string `gorm:"varchar(20);not null"`
	Email     string
	Birthday  string `gorm:"default:20000101"`
	Signature string `gorm:"default:This man is lazy."`
	gorm.Model
}

const PwdHard = 12

func SetPwd(pwd string) string {
	pwdByte, _ := bcrypt.GenerateFromPassword([]byte(pwd), PwdHard)
	return string(pwdByte)
}

func (receiver *User) CheckPwd(pwd string) bool {
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
