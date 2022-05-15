package dao

import (
	"log"
	"task/model"
)

func CreateTask(task model.Task) (err error) {
	err = model.DB.Model(&model.Task{}).Create(&task).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func FindListTask(status int64, id uint64) (taskList []model.Task, err error) {
	err = model.DB.Model(&model.Task{}).Where("status = ? and uid = ?", status, id).
		Find(&taskList).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func FindOneTask(content string) (task []model.Task, err error) {
	err = model.DB.Model(&model.Task{}).
		Where("title like ? or content like ?", "%"+content+"%", "%"+content+"%").
		Find(&task).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func UpdateTask(id uint64, uid uint64) (task model.Task, err error) {
	err = model.DB.Model(&model.Task{}).Where("id = ? and uid = ?", id, uid).First(&task).Error
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

func DeleteTask(id uint64, uid uint64) (err error) {
	var task model.Task
	err = model.DB.Model(&model.Task{}).Where("id = ? and uid = ?", id, uid).Delete(&task).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}
