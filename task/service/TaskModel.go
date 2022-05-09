package service

import (
	"context"
	"gorm.io/gorm"
	"log"
	"task/dao"
	"task/model"
	services "task/proto/pb"
)

type UserTask struct {
}

var T = &UserTask{}

func (u *UserTask) CreateTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskResponse) error {
	var task model.Task
	task = model.Task{
		Model:   gorm.Model{},
		Id:      0,
		Title:   req.Title,
		Content: req.Content,
	}
	err := dao.CreateTask(task)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
