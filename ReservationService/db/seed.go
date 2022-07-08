package db

import (
	"ReservationService/models"
	"gorm.io/gorm"
)

var Reservations = []models.Reservation{
	{
		Model:             gorm.Model{},
		DriveId:           1,
		PassengerUsername: "beli",
		DriverUsername:    "tica",
		Verified:          false,
	},
	{
		Model:             gorm.Model{},
		DriveId:           2,
		PassengerUsername: "beli",
		DriverUsername:    "tica",
		Verified:          true,
	},
	{
		Model:             gorm.Model{},
		DriveId:           2,
		PassengerUsername: "boki",
		DriverUsername:    "tica",
		Verified:          true,
	},
	{
		Model:             gorm.Model{},
		DriveId:           3,
		PassengerUsername: "boki",
		DriverUsername:    "tica",
		Verified:          true,
	},
}
