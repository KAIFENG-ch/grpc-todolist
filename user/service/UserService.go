package service

import (
	"context"
	"errors"
	"time"
	"user/dao"
	"user/model"
	services "user/proto/proto"
	"user/utils"
)

type UserService struct {
}

var U = UserService{}

func (us *UserService) UserRegister(ctx context.Context, req *services.UserRequest) (resp *services.UserDetailResponse, err error) {
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码不一致")
		return nil, err
	}
	_, ok := dao.FindUser(req.UserName)
	if ok {
		resp = &services.UserDetailResponse{Message: "用户名已存在", Code: 500}
		return resp, nil
	}
	userInfo := model.User{
		Username: req.UserName,
		Password: req.Password,
	}
	err = dao.CreateUser(userInfo)
	if err != nil {
		return nil, err
	}
	resp = &services.UserDetailResponse{
		UserDetail: &services.UserModel{
			UserName:  userInfo.Username,
			CreatedAt: time.Now().Unix(),
		},
		Message: "创建成功！",
		Code:    200,
	}
	return resp, nil
}

func (us *UserService) UserLogin(ctx context.Context, req *services.UserRequest) (resp *services.UserDetailResponse, err error) {
	resp = new(services.UserDetailResponse)
	user, ok := dao.FindUser(req.UserName)
	if !ok {
		err = errors.New("用户不存在")
		resp.Code = 400
		return
	}
	if user.CheckPwd(req.Password) {
		resp.Code = 400
		return
	}
	token, err := utils.GenerateToken(user.ID)
	resp.UserDetail = &services.UserModel{
		UserName: user.Username,
	}
	resp.Message = "token: " + token
	resp.Code = 200
	return
}
