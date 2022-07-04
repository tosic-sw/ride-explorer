package router

import (
	"UserService/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func MapRoutesAndServe(handler *handlers.UsersHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/users/login", handler.Login).Methods("POST")
	router.HandleFunc("/api/users/passenger-registration", handler.AdminRegistration).Methods("POST")
	router.HandleFunc("/api/users/driver-registration", handler.DriverRegistration).Methods("POST")
	router.HandleFunc("/api/users/passenger-registration", handler.PassengerRegistration).Methods("POST")

	// Izmeni - mozda i ne
	router.HandleFunc("/api/users/search-admin", handler.SearchAdmin).Methods("GET")
	router.HandleFunc("/api/users/search-driver", handler.SearchDriver).Methods("GET")
	router.HandleFunc("/api/users/search-passenger", handler.SearchPassenger).Methods("GET")

	router.HandleFunc("/api/users/admin/{username}", handler.GetAdmin).Methods("GET")
	router.HandleFunc("/api/users/driver/{username}", handler.GetDriver).Methods("GET")
	router.HandleFunc("/api/users/passenger/{username}", handler.GetPassenger).Methods("GET")

	router.HandleFunc("/api/users/passenger/{username}", handler.UpdateUserData).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", router))
}
