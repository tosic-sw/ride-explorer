package main

import (
	"UserService/data"
	"UserService/db"
	"UserService/handlers"
	"UserService/router"
)

func main() {
	dbConn := db.Init()

	repository := data.NewRepository(dbConn)

	usersHandler := handlers.NewUsersHandler(repository)

	router.MapRoutesAndServe(usersHandler)
}
