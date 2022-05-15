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
	resp = new(services.TaskResponse)
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

func (u *UserTask) GetListTask(ctx context.Context, req *services.TaskRequest) (resp *services.TaskListResponse, err error) {
	resp = new(services.TaskListResponse)
	tasks, err := dao.FindListTask(req.Status, req.Uid)
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

func (u *UserTask) GetSomeTask(ctx context.Context, req *services.TaskRequest) (resp *services.TaskListResponse, err error) {
	resp = new(services.TaskListResponse)
	tasks, err := dao.FindOneTask(req.Content)
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

func (u *UserTask) UpdateTask(ctx context.Context, req *services.TaskRequest) (resp *services.TaskResponse, err error) {
	resp = new(services.TaskResponse)
	task, err := dao.UpdateTask(req.Id, req.Uid)
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

func (u *UserTask) DeleteTask(ctx context.Context, req *services.TaskRequest) (resp *services.TaskResponse, err error) {
	resp = new(services.TaskResponse)
	resp.Task = new(services.TaskModel)
	err = dao.DeleteTask(req.Id, req.Uid)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resp.Task.Id = req.Id
	return
}
