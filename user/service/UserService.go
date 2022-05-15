package service

import (
	"context"
	"errors"
	"time"
	"user/dao"
	"user/model"
	services "user/proto/pb"
)

type UserService struct {
}

var U = UserService{}

func (us *UserService) UserRegister(ctx context.Context, req *services.UserRequest) (resp *services.UserDetailResponse, err error) {
	resp = new(services.UserDetailResponse)
	_, ok := dao.FindUser(req.UserName)
	if ok {
		err = errors.New("用户名已存在")
		return
	}
	userInfo := model.User{
		Username: req.UserName,
		Password: model.SetPwd(req.Password),
	}
	err = dao.CreateUser(userInfo)
	if err != nil {
		return nil, err
	}
	resp.UserDetail = &services.UserModel{
		UserName:  userInfo.Username,
		CreatedAt: time.Now().Unix(),
	}
	resp.Message = "创建成功！"
	resp.Code = 200
	return
}

func (us *UserService) UserLogin(ctx context.Context, req *services.UserRequest) (resp *services.UserDetailResponse, err error) {
	resp = new(services.UserDetailResponse)
	user, ok := dao.FindUser(req.UserName)
	if !ok {
		err = errors.New("用户不存在")
		resp.Code = 400
		return
	}
	if !user.CheckPwd(req.Password) {
		resp.Code = 400
		return
	}
	if err != nil {
		panic(err)
	}
	resp.UserDetail = &services.UserModel{
		Id:       uint32(user.ID),
		UserName: user.Username,
	}
	resp.Message = "登录成功！"
	resp.Code = 200
	return
}
