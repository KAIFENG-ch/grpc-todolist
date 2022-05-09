package dao

import (
	"log"
	"task/model"
)

func CreateTask(task model.Task) (err error) {
	err = model.DB.Model(&task).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return
}
