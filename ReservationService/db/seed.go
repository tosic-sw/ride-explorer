package db

import (
	"ReservationService/models"
	"gorm.io/gorm"
)

var Reservations = []models.Reservation{
	{
		Model:             gorm.Model{},
		DriveId:           1, // Tica od Beceja do Novog Sada
		PassengerUsername: "beli",
		DriverUsername:    "tica",
		Verified:          false,
	},
	{
		Model:             gorm.Model{},
		DriveId:           2, // Ukica od Zajecara do Novog Sada
		PassengerUsername: "toma",
		DriverUsername:    "ukica",
		Verified:          true,
	},
	{
		Model:             gorm.Model{},
		DriveId:           2, // Ukica od Zajecara do Novog Sada
		PassengerUsername: "mile",
		DriverUsername:    "ukica",
		Verified:          false,
	},
	{
		Model:             gorm.Model{},
		DriveId:           3, // Ukica od Novog Sada do Beograda
		PassengerUsername: "mile",
		DriverUsername:    "ukica",
		Verified:          false,
	},
	{
		Model:             gorm.Model{},
		DriveId:           4, // Ukica od Novog Sada do Beograda u proslosti da bi mogao finish
		PassengerUsername: "toma",
		DriverUsername:    "ukica",
		Verified:          true,
	},
	{
		Model:             gorm.Model{},
		DriveId:           4, // Ukica od Novog Sada do Beograda u prosloti da bi mogao finish
		PassengerUsername: "mile",
		DriverUsername:    "ukica",
		Verified:          true,
	},
}
