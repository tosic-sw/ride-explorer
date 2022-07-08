package router

import (
	"ReservationService/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func MapRoutesAndServe(handler *handlers.ReservationsHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/reservations/{id:[0-9]+}", handler.GetReservation).Methods(http.MethodGet)
	router.HandleFunc("/api/reservations", handler.CreateReservation).Methods(http.MethodPost)
	router.HandleFunc("/api/reservations/{id:[0-9]+}", handler.DeleteReservation).Methods(http.MethodDelete)
	router.HandleFunc("/api/reservations/verify/{id:[0-9]+}", handler.VerifyReservation).Methods(http.MethodPut)
	router.HandleFunc("/api/reservations/user/verified", handler.GetAllByUserVerified).Methods(http.MethodGet)
	router.HandleFunc("/api/reservations/user/unverified", handler.GetAllByUserUnverified).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8082", router))
}
