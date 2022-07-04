package db

import (
	"UserService/models"
	"gorm.io/gorm"
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
			Password: "admin",
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
		Banned:    false,
		UserAccount: models.UserAccount{
			Model:    gorm.Model{},
			Password: "beli",
		},
	},
	{
		Model:     gorm.Model{},
		Email:     "boki@maildrop.cc",
		Username:  "boki",
		Firstname: "Bojan",
		Lastname:  "Baskalo",
		Role:      "PASSENGER",
		Banned:    true,
		UserAccount: models.UserAccount{
			Model:    gorm.Model{},
			Password: "boki",
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
		UserAccount: models.UserAccount{
			Model:    gorm.Model{},
			Password: "tica",
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
		Banned: false,
	},
	{
		Model:     gorm.Model{},
		Email:     "ukica@maildrop.cc",
		Username:  "ukica",
		Firstname: "Uros",
		Lastname:  "Stojanovic",
		Role:      "DRIVER",
		UserAccount: models.UserAccount{
			Model:    gorm.Model{},
			Password: "ukica",
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
		Banned: false,
	},
}

// #################################### //
