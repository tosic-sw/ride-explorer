package models

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model
	DriveId           int32  `gorm:"not null"`
	PassengerUsername string `gorm:"not null"`
	DriverUsername    string `gorm:"not null"`
	Verified          bool   `gorm:"not null"`
}
