package models

import "gorm.io/gorm"

type UserAccount struct {
	gorm.Model
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	Role     Role   `gorm:"not null"`
	UserID   int
	UserType string
}

type Admin struct {
	gorm.Model
	Email       string      `gorm:"not null;unique"`
	Username    string      `gorm:"not null;unique"`
	Firstname   string      `gorm:"not null"`
	Lastname    string      `gorm:"not null"`
	Role        Role        `gorm:"not null"`
	UserAccount UserAccount `gorm:"polymorphic:User;"`
}

type Driver struct {
	gorm.Model
	Email       string      `gorm:"not null;unique"`
	Username    string      `gorm:"not null;unique"`
	Firstname   string      `gorm:"not null"`
	Lastname    string      `gorm:"not null"`
	Role        Role        `gorm:"not null"`
	UserAccount UserAccount `gorm:"polymorphic:User;"`
	Car         Car
	Banned      bool
}

type Passenger struct {
	gorm.Model
	Email       string      `gorm:"not null;unique"`
	Username    string      `gorm:"not null;unique"`
	Firstname   string      `gorm:"not null"`
	Lastname    string      `gorm:"not null"`
	Role        Role        `gorm:"not null"`
	UserAccount UserAccount `gorm:"polymorphic:User;"`
	Banned      bool
}

type Car struct {
	gorm.Model
	PlateNumber     string  `gorm:"not null;unique"`
	Brand           string  `gorm:"not null"`
	CarModel        string  `gorm:"not null"`
	FuelConsumption float32 `gorm:"not null"`
	Volume          float32 `gorm:"not null"`
	Power           float32 `gorm:"not null"`
	DriverId        uint
}
