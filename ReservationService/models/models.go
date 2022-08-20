package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	DriveId           int32  `gorm:"not null"`
	PassengerUsername string `gorm:"not null"`
	DriverUsername    string `gorm:"not null"`
	Verified          bool   `gorm:"not null"`
}

func (reservation *Reservation) GetPassengerEmail() string {
	return fmt.Sprintf("%s@maildrop.cc", reservation.PassengerUsername)
}

func (reservation *Reservation) GetDriverEmail() string {
	return fmt.Sprintf("%s@maildrop.cc", reservation.DriverUsername)
}

type Mail struct {
	To      string
	Subject string
	Body    string
}
