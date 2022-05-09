package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Id      uint   `gorm:"not null"`
	Title   string `gorm:"varchar(30),not null"`
	Content string
}
