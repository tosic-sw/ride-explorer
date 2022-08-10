package db

import (
	"UserService/models"
	"gorm.io/gorm"
	"time"
)

// #################################### //

var Admins = []models.Admin{
	{
		Model:     gorm.Model{},
		Email:     "veljko@maildrop.cc",
		Username:  "admin",
		Firstname: "Veljko",
		Lastname:  "Tosic",
		Role:      "ADMIN",
		UserAccount: models.UserAccount{
			Model:    gorm.Model{},
			Username: "admin",
			Password: "admin",
			Role:     models.ADMIN,
			Verified: true,
		},
	},
	{
		Model:     gorm.Model{},
		Email:     "milic@maildrop.cc",
		Username:  "milic",
		Firstname: "Dragan",
		Lastname:  "Milic",
		Role:      "ADMIN",
		UserAccount: models.UserAccount{
			Model:    gorm.Model{},
			Username: "milic",
			Password: "milic",
			Role:     models.ADMIN,
			Verified: true,
		},
	},
}

// #################################### //

var Passengers = []models.Passenger{
	{
		Model:     gorm.Model{},
		Email:     "beli@maildrop.cc",
		Username:  "beli",
		Firstname: "Marko",
		Lastname:  "Bjelica",
		Role:      "PASSENGER",
		UserAccount: models.UserAccount{
			Model:    gorm.Model{},
			Username: "beli",
			Password: "beli",
			Role:     models.PASSENGER,
			Verified: true,
		},
	},
	{
		Model:       gorm.Model{},
		Email:       "boki@maildrop.cc",
		Username:    "boki",
		Firstname:   "Bojan",
		Lastname:    "Baskalo",
		Role:        "PASSENGER",
		BannedUntil: time.Now().AddDate(0, 3, 0).UnixMilli(),
		UserAccount: models.UserAccount{
			Model:       gorm.Model{},
			Username:    "boki",
			Password:    "boki",
			Role:        models.PASSENGER,
			BannedUntil: time.Now().AddDate(0, 3, 0).UnixMilli(),
			Verified:    true,
		},
	},
}

// #################################### //

var Drivers = []models.Driver{
	{
		Model:     gorm.Model{},
		Email:     "tica@maildrop.cc",
		Username:  "tica",
		Firstname: "Darko",
		Lastname:  "Tica",
		Role:      "DRIVER",
		Verified:  true,
		UserAccount: models.UserAccount{
			Model:    gorm.Model{},
			Username: "tica",
			Password: "tica",
			Role:     models.DRIVER,
			Verified: true,
		},
		Car: models.Car{
			Model:           gorm.Model{},
			PlateNumber:     "NS_BOSS",
			Brand:           "Golf",
			CarModel:        "8",
			FuelConsumption: 10,
			Volume:          1600,
			Power:           105,
		},
	},
	{
		Model:     gorm.Model{},
		Email:     "ukica@maildrop.cc",
		Username:  "ukica",
		Firstname: "Uros",
		Lastname:  "Stojanovic",
		Role:      "DRIVER",
		Verified:  true,
		UserAccount: models.UserAccount{
			Model:    gorm.Model{},
			Username: "ukica",
			Password: "ukica",
			Role:     models.DRIVER,
			Verified: true,
		},
		Car: models.Car{
			Model:           gorm.Model{},
			PlateNumber:     "NS_OOP",
			Brand:           "Peugeot",
			CarModel:        "308",
			FuelConsumption: 12,
			Volume:          1600,
			Power:           101,
		},
	},
	{
		Model:     gorm.Model{},
		Email:     "saomi@maildrop.cc",
		Username:  "saomi",
		Firstname: "Milos",
		Lastname:  "Manojlovic",
		Role:      "DRIVER",
		Verified:  false,
		UserAccount: models.UserAccount{
			Model:    gorm.Model{},
			Username: "saomi",
			Password: "saomi",
			Role:     models.DRIVER,
			Verified: false,
		},
		Car: models.Car{
			Model:           gorm.Model{},
			PlateNumber:     "740TC",
			Brand:           "Audi",
			CarModel:        "A3",
			FuelConsumption: 10,
			Volume:          1600,
			Power:           105,
		},
	},
}

// #################################### //
