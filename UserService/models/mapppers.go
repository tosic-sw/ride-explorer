package models

import "gorm.io/gorm"

func (admin *Admin) ToDTO() UserDTO {
	return UserDTO{
		Email:     admin.Email,
		Username:  admin.Username,
		Firstname: admin.Firstname,
		Lastname:  admin.Lastname,
		Role:      string(admin.Role),
	}
}

func (admin *Admin) ToUpdateDTO() UserForUpdateDTO {
	return UserForUpdateDTO{
		Email:     admin.Email,
		Firstname: admin.Firstname,
		Lastname:  admin.Lastname,
	}
}

func (regDTO *RegistrationDTO) ToAdmin() *Admin {
	return &Admin{
		Model:     gorm.Model{},
		Email:     regDTO.Email,
		Username:  regDTO.Username,
		Firstname: regDTO.Firstname,
		Lastname:  regDTO.Lastname,
		Role:      ADMIN,
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
		Email:     driver.Email,
		Username:  driver.Username,
		Firstname: driver.Firstname,
		Lastname:  driver.Lastname,
		Role:      string(driver.Role),
	}
}

func (driver *Driver) ToUpdateDTO() UserForUpdateDTO {
	return UserForUpdateDTO{
		Email:     driver.Email,
		Firstname: driver.Firstname,
		Lastname:  driver.Lastname,
	}
}

func (regDTO *DriverRegistrationDTO) ToDriver() *Driver {
	return &Driver{
		Model:     gorm.Model{},
		Email:     regDTO.Email,
		Username:  regDTO.Username,
		Firstname: regDTO.Firstname,
		Lastname:  regDTO.Lastname,
		Role:      DRIVER,
		Verified:  false,
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
		Email:     passenger.Email,
		Username:  passenger.Username,
		Firstname: passenger.Firstname,
		Lastname:  passenger.Lastname,
		Role:      string(passenger.Role),
	}
}

func (passenger *Passenger) ToUpdateDTO() UserForUpdateDTO {
	return UserForUpdateDTO{
		Email:     passenger.Email,
		Firstname: passenger.Firstname,
		Lastname:  passenger.Lastname,
	}
}

func (regDTO *RegistrationDTO) ToPassenger() *Passenger {
	return &Passenger{
		Model:     gorm.Model{},
		Email:     regDTO.Email,
		Username:  regDTO.Username,
		Firstname: regDTO.Firstname,
		Lastname:  regDTO.Lastname,
		Role:      PASSENGER,
		UserAccount: UserAccount{
			Model:    gorm.Model{},
			Username: regDTO.Username,
			Password: regDTO.Password,
			Role:     PASSENGER,
			Verified: true,
		},
	}
}
