package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	Evaluator string `gorm:"not null"`
	Evaluated string `gorm:"not null"`
	DriveId   uint   `gorm:"not null"`
	Positive  bool   `gorm:"not null"`
	Text      string `gorm:"not null"`
}
