package models

import "gorm.io/gorm"

type Complaint struct {
	gorm.Model
	Accuser string `gorm:"not null"`
	Accused string `gorm:"not null"`
	DriveId uint   `gorm:"not null"`
	Text    string `gorm:"not null"`
}
