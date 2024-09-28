package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Status string `gorm:"index:status_idx"`
}
