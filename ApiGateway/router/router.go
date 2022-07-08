package router

import (
	"ApiGateway/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func MapRoutesAndServe() {
	router := mux.NewRouter()

	router.HandleFunc("/api/users/login", handlers.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/users/registration/admin", handlers.RegistrationAdmin).Methods(http.MethodPost)
	router.HandleFunc("/api/users/registration/driver", handlers.RegistrationDriver).Methods(http.MethodPost)
	router.HandleFunc("/api/users/registration/driver/verify/{username}", handlers.DriverVerification).Methods(http.MethodPut)
	router.HandleFunc("/api/users/registration/passenger", handlers.RegistrationPassenger).Methods(http.MethodPost)
	router.HandleFunc("/api/users/search/admin", handlers.SearchAdmin).Methods(http.MethodGet)
	router.HandleFunc("/api/users/search/driver", handlers.SearchDriver).Methods(http.MethodGet)
	router.HandleFunc("/api/users/search/passenger", handlers.SearchPassenger).Methods(http.MethodGet)
	router.HandleFunc("/api/users/admin/{username}", handlers.GetAdmin).Methods(http.MethodGet)
	router.HandleFunc("/api/users/driver/{username}", handlers.GetDriver).Methods(http.MethodGet)
	router.HandleFunc("/api/users/passenger/{username}", handlers.GetPassenger).Methods(http.MethodGet)
	router.HandleFunc("/api/users/profile", handlers.UpdateProfile).Methods(http.MethodPut)
	router.HandleFunc("/api/users/change-password", handlers.ChangePassword).Methods(http.MethodPut)
	router.HandleFunc("/api/users/ban/driver/{username}", handlers.BanDriver).Methods(http.MethodPut)
	router.HandleFunc("/api/users/ban/passenger/{username}", handlers.BanPassenger).Methods(http.MethodPut)
	router.HandleFunc("/api/users/driver/{username}", handlers.DeleteDriver).Methods(http.MethodDelete)
	router.HandleFunc("/api/users/passenger/{username}", handlers.DeletePassenger).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", router))
}