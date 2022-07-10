package main

import (
	"RatingService/data"
	"RatingService/db"
	"RatingService/handlers"
	"RatingService/router"
)

func main() {
	dbConn := db.Init()

	repository := data.NewRepository(dbConn)

	ratingsHandler := handlers.NewRatingsHandler(repository)

	router.MapRoutesAndServe(ratingsHandler)
}
