package dao

import (
	"log"
	"task/model"
)

func CreateTask(task model.Task) (err error) {
	err = model.DB.Model(&task).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func FindListTask(status int64) (taskList []model.Task, err error) {
	err = model.DB.Model(&model.Task{}).Where("status = ?", status).Find(&taskList).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func FindOneTask(id uint64) (task model.Task, err error) {
	err = model.DB.Model(&model.Task{}).Where("id = ?", id).First(&task).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func UpdateTask(id uint64) (task model.Task, err error) {
	err = model.DB.Model(&model.Task{}).Where("id = ?", id).First(&task).Error
	if err != nil {
		log.Println(err)
		return
	}
	if task.Status == 0 {
		model.DB.Model(&task).Update("status", 1)
	} else {
		model.DB.Model(&task).Update("status", 0)
	}
	return
}

func DeleteTask(id uint64) (err error) {
	var task model.Task
	err = model.DB.Model(&model.Task{}).Where("id = ?", id).Delete(&task).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}
