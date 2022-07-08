package models

type ReservationDTO struct {
	Id                uint   `json:"id"`
	CreatedAt         int64  `json:"createdAt"`
	DriveId           int32  `json:"driveId"`
	PassengerUsername string `json:"passengerUsername"`
	DriverUsername    string `json:"driverUsername"`
	Verified          bool   `json:"verified"`
}

type CreateReservationDTO struct {
	DriveId        int32  `json:"driveId"`
	DriverUsername string `json:"driverUsername"`
}

type Response struct {
	Message string `json:"message"`
}

type ReserveDTO struct {
	Id     int32 `json:"id"`
	Places int32 `json:"places"`
}
