package main

import (
	"ComplaintService/data"
	"ComplaintService/db"
	"ComplaintService/handlers"
	"ComplaintService/router"
)

func main() {
	dbConn := db.Init()

	repository := data.NewRepository(dbConn)

	complaintsHandler := handlers.NewComplaintsHandler(repository)

	router.MapRoutesAndServe(complaintsHandler)
}
