package main

import (
	"ReservationService/data"
	"ReservationService/db"
	"ReservationService/handlers"
	"ReservationService/router"
)

func main() {
	dbConn := db.Init()

	repository := data.NewRepository(dbConn)

	reservationsHandler := handlers.NewReservationsHandler(repository)

	router.MapRoutesAndServe(reservationsHandler)
}
