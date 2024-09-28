package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Status  string `gorm:"default:'pending';index:status_idx"`
	To      string
	Subject string
	Body    string
}
