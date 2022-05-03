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

func (us *UserService) UserRegister(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码不一致")
		return err
	}
	userInfo,err := dao.FindUser(req.UserName)
	if err != nil {
		err = errors.New("用户名已存在")
	}
	err = dao.CreateUser(userInfo)
	if err != nil {
		return err
	}
	resp.UserDetail = model.Build(userInfo)
	return nil
}

func (us *UserService) LoginRequest(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	users, err := dao.FindUser(req.UserName)
	if err != nil {
		return nil
	}
	if users.CheckPwd(req.Password) {
		resp.Code = 400
		return nil
	}
	resp.UserDetail = model.Build(users)
	return nil
}