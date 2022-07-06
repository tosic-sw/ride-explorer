package models

type UserAccountDTO struct {
	Password string `json:"role"`
}

type UserDTO struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Role      string `json:"role"`
}

type RegistrationDTO struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
}

type DriverRegistrationDTO struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Car       CarDTO `json:"car"`
}

type UserForUpdateDTO struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type UserPasswordChangeDTO struct {
	Password string `json:"password"`
}

type DriverDTO struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Role      string `json:"role"`
}

type DriverWithCarDTO struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Role      string `json:"role"`
	Car       CarDTO `json:"car"`
}

type CarDTO struct {
	PlateNumber     string  `json:"plateNumber"`
	Brand           string  `json:"brand"`
	CarModel        string  `json:"carModel"`
	FuelConsumption float32 `json:"fuelConsumption"`
	Volume          float32 `json:"volume"`
	Power           float32 `json:"power"`
}

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type TokenState struct {
	Token string `json:"token"`
}
