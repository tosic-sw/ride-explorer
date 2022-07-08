package models

import "gorm.io/gorm"

func (reservation *Reservation) ToDTO() ReservationDTO {
	return ReservationDTO{
		Id:                reservation.ID,
		CreatedAt:         reservation.CreatedAt.UnixMilli(),
		DriveId:           reservation.DriveId,
		PassengerUsername: reservation.PassengerUsername,
		Verified:          reservation.Verified,
		DriverUsername:    reservation.DriverUsername,
	}
}

func (dto *CreateReservationDTO) ToReservation(passengerUsername string) *Reservation {
	return &Reservation{
		Model:             gorm.Model{},
		DriveId:           dto.DriveId,
		PassengerUsername: passengerUsername,
		DriverUsername:    dto.DriverUsername,
		Verified:          false,
	}
}
