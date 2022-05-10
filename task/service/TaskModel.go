package service

import (
	"context"
	"log"
	"task/dao"
	"task/model"
	services "task/proto/pb"
)

type UserTask struct {
}

var T = UserTask{}

func (u *UserTask) CreateTask(ctx context.Context, req *services.TaskRequest) (resp *services.TaskResponse, err error) {
	var task model.Task
	task = model.Task{
		Uid:     req.Uid,
		Status:  0,
		Title:   req.Title,
		Content: req.Content,
	}
	err = dao.CreateTask(task)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resp.Task = &services.TaskModel{
		Id:        uint64(task.ID),
		Status:    task.Status,
		Title:     task.Title,
		Content:   task.Content,
		Uid:       task.Uid,
		StartAt:   task.CreatedAt.Unix(),
		UpdatedAt: task.UpdatedAt.Unix(),
		EndAt:     0,
	}
	return
}

func (u UserTask) GetListTask(ctx context.Context, req *services.TaskRequest) (resp *services.TaskListResponse, err error) {
	tasks, err := dao.FindListTask(req.Status)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var taskModels []*services.TaskModel
	for _, t := range tasks {
		task := &services.TaskModel{
			Id:        uint64(t.ID),
			Uid:       t.Uid,
			Title:     t.Title,
			Content:   t.Content,
			StartAt:   t.CreatedAt.Unix(),
			Status:    t.Status,
			UpdatedAt: t.UpdatedAt.Unix(),
			EndAt:     0,
		}
		taskModels = append(taskModels, task)
	}
	resp.TaskList = taskModels
	resp.Count = uint32(len(taskModels))
	return
}

func (u UserTask) GetOneTask(ctx context.Context, req *services.TaskRequest) (resp *services.TaskResponse, err error) {
	task, err := dao.FindOneTask(req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resp.Task = &services.TaskModel{
		Id:        uint64(task.ID),
		Uid:       task.Uid,
		Title:     task.Title,
		Content:   task.Content,
		Status:    task.Status,
		StartAt:   task.CreatedAt.Unix(),
		UpdatedAt: task.UpdatedAt.Unix(),
		EndAt:     0,
	}
	return
}

func (u UserTask) UpdateTask(ctx context.Context, req *services.TaskRequest) (resp *services.TaskResponse, err error) {
	task, err := dao.UpdateTask(req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resp.Task = &services.TaskModel{
		Id:        uint64(task.ID),
		Uid:       task.Uid,
		Title:     task.Title,
		Content:   task.Content,
		StartAt:   task.CreatedAt.Unix(),
		EndAt:     0,
		Status:    task.Status,
		UpdatedAt: task.UpdatedAt.Unix(),
	}
	return
}

func (u UserTask) DeleteTask(ctx context.Context, req *services.TaskRequest) (resp *services.TaskResponse, err error) {
	err = dao.DeleteTask(req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resp.Task = &services.TaskModel{}
	return
}
