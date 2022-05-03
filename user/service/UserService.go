package service

import (
	"context"
	"errors"
	"user/dao"
	"user/model"
	services "user/proto"
)

type UserService struct {
}

var U = UserService{}

func (us *UserService) UserRegister(ctx context.Context, req *services.UserRequest) (resp *services.UserDetailResponse, err error) {
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码不一致")
		return nil, err
	}
	userInfo, err := dao.FindUser(req.UserName)
	if err != nil {
		err = errors.New("用户名已存在")
	}
	err = dao.CreateUser(userInfo)
	if err != nil {
		return nil, err
	}
	resp.UserDetail = model.Build(userInfo)
	resp.Code = 200
	return
}

func (us *UserService) UserLogin(ctx context.Context, req *services.UserRequest) (resp *services.UserDetailResponse, err error) {
	users, err := dao.FindUser(req.UserName)
	if err != nil {
		resp.Code = 400
		return
	}
	if users.CheckPwd(req.Password) {
		resp.Code = 400
		return
	}
	resp.UserDetail = model.Build(users)
	resp.Code = 200
	return
}
