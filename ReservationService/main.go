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

	mailHandler := handlers.NewMailHandler()

	reservationsHandler := handlers.NewReservationsHandler(repository, mailHandler)

	router.MapRoutesAndServe(reservationsHandler)
}
