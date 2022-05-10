package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Uid     uint64 `gorm:"not null"`
	Status  int64  `gorm:"default:0"`
	Title   string `gorm:"varchar(30),not null"`
	Content string
}
