package models

import "gorm.io/gorm"

type UserAccount struct {
	gorm.Model
	Password string `gorm:"not null"`
	UserID   int
	UserType string
}

// ################ Admin ###################

type Admin struct {
	gorm.Model
	Email       string      `gorm:"not null;unique"`
	Username    string      `gorm:"not null;unique"`
	Firstname   string      `gorm:"not null"`
	Lastname    string      `gorm:"not null"`
	Role        Role        `gorm:"not null"`
	UserAccount UserAccount `gorm:"polymorphic:User;"`
}

func (admin *Admin) ToDTO() UserDTO {
	return UserDTO{
		Email:     admin.Email,
		Username:  admin.Username,
		Firstname: admin.Firstname,
		Lastname:  admin.Lastname,
		Role:      string(admin.Role),
	}
}

// ################ Driver ###################

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

func (admin *Driver) ToDTO() UserDTO {
	return UserDTO{
		Email:     admin.Email,
		Username:  admin.Username,
		Firstname: admin.Firstname,
		Lastname:  admin.Lastname,
		Role:      string(admin.Role),
	}
}

// ################ Passenger ###################

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

func (admin *Passenger) ToDTO() UserDTO {
	return UserDTO{
		Email:     admin.Email,
		Username:  admin.Username,
		Firstname: admin.Firstname,
		Lastname:  admin.Lastname,
		Role:      string(admin.Role),
	}
}

// ################ Car ###################

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
