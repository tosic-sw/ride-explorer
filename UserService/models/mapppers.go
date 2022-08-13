package models

import "gorm.io/gorm"

func (admin *Admin) ToDTO() UserDTO {
	return UserDTO{
		Email:       admin.Email,
		Username:    admin.Username,
		Firstname:   admin.Firstname,
		Lastname:    admin.Lastname,
		PhoneNumber: admin.PhoneNumber,
		Role:        string(admin.Role),
	}
}

func (admin *Admin) ToUpdateDTO() UserForUpdateDTO {
	return UserForUpdateDTO{
		Email:       admin.Email,
		Password:    admin.UserAccount.Password,
		Firstname:   admin.Firstname,
		Lastname:    admin.Lastname,
		PhoneNumber: admin.PhoneNumber,
	}
}

func (regDTO *RegistrationDTO) ToAdmin() *Admin {
	return &Admin{
		Model:       gorm.Model{},
		Email:       regDTO.Email,
		Username:    regDTO.Username,
		Firstname:   regDTO.Firstname,
		Lastname:    regDTO.Lastname,
		PhoneNumber: regDTO.PhoneNumber,
		Role:        ADMIN,
		UserAccount: UserAccount{
			Model:    gorm.Model{},
			Username: regDTO.Username,
			Password: regDTO.Password,
			Role:     ADMIN,
			Verified: true,
		},
	}
}

func (driver *Driver) ToDTO() UserDTO {
	return UserDTO{
		Email:       driver.Email,
		Username:    driver.Username,
		Firstname:   driver.Firstname,
		Lastname:    driver.Lastname,
		PhoneNumber: driver.PhoneNumber,
		Role:        string(driver.Role),
	}
}

func (driver *Driver) ToUpdateDTO() UserForUpdateDTO {
	return UserForUpdateDTO{
		Email:       driver.Email,
		Password:    driver.UserAccount.Password,
		Firstname:   driver.Firstname,
		Lastname:    driver.Lastname,
		PhoneNumber: driver.PhoneNumber,
	}
}

func (driver *Driver) ToDriverDTO() DriverWithCarDTO {
	return DriverWithCarDTO{
		Email:       driver.Email,
		Username:    driver.Username,
		Firstname:   driver.Firstname,
		Lastname:    driver.Lastname,
		PhoneNumber: driver.PhoneNumber,
		Role:        string(driver.Role),
		Car: CarDTO{
			PlateNumber:     driver.Car.PlateNumber,
			Brand:           driver.Car.Brand,
			CarModel:        driver.Car.CarModel,
			FuelConsumption: driver.Car.FuelConsumption,
			Volume:          driver.Car.Volume,
			Power:           driver.Car.Power,
		},
	}
}

func (regDTO *DriverRegistrationDTO) ToDriver() *Driver {
	return &Driver{
		Model:       gorm.Model{},
		Email:       regDTO.Email,
		Username:    regDTO.Username,
		Firstname:   regDTO.Firstname,
		Lastname:    regDTO.Lastname,
		PhoneNumber: regDTO.PhoneNumber,
		Role:        DRIVER,
		Verified:    false,
		UserAccount: UserAccount{
			Model:    gorm.Model{},
			Username: regDTO.Username,
			Password: regDTO.Password,
			Role:     DRIVER,
			Verified: false,
		},
		Car: Car{
			Model:           gorm.Model{},
			PlateNumber:     regDTO.Car.PlateNumber,
			Brand:           regDTO.Car.Brand,
			CarModel:        regDTO.Car.CarModel,
			FuelConsumption: regDTO.Car.FuelConsumption,
			Volume:          regDTO.Car.Volume,
			Power:           regDTO.Car.Power,
		},
	}
}

func (passenger *Passenger) ToDTO() UserDTO {
	return UserDTO{
		Email:       passenger.Email,
		Username:    passenger.Username,
		Firstname:   passenger.Firstname,
		Lastname:    passenger.Lastname,
		PhoneNumber: passenger.PhoneNumber,
		Role:        string(passenger.Role),
	}
}

func (passenger *Passenger) ToUpdateDTO() UserForUpdateDTO {
	return UserForUpdateDTO{
		Email:       passenger.Email,
		Password:    passenger.UserAccount.Password,
		Firstname:   passenger.Firstname,
		Lastname:    passenger.Lastname,
		PhoneNumber: passenger.PhoneNumber,
	}
}

func (regDTO *RegistrationDTO) ToPassenger() *Passenger {
	return &Passenger{
		Model:       gorm.Model{},
		Email:       regDTO.Email,
		Username:    regDTO.Username,
		Firstname:   regDTO.Firstname,
		Lastname:    regDTO.Lastname,
		PhoneNumber: regDTO.PhoneNumber,
		Role:        PASSENGER,
		UserAccount: UserAccount{
			Model:    gorm.Model{},
			Username: regDTO.Username,
			Password: regDTO.Password,
			Role:     PASSENGER,
			Verified: true,
		},
	}
}
