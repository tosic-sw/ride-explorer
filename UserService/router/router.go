package router

import (
	"UserService/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func MapRoutesAndServe(handler *handlers.UsersHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/users/login", handler.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/users/authenticate", handler.Authenticate).Methods(http.MethodGet)
	router.HandleFunc("/api/users/authorize/admin", handler.AuthorizeAdmin).Methods(http.MethodGet)
	router.HandleFunc("/api/users/authorize/driver", handler.AuthorizeDriver).Methods(http.MethodGet)
	router.HandleFunc("/api/users/authorize/passenger", handler.AuthorizePassenger).Methods(http.MethodGet)

	router.HandleFunc("/api/users/registration/admin", handler.AdminRegistration).Methods(http.MethodPost)
	router.HandleFunc("/api/users/registration/driver", handler.DriverRegistration).Methods(http.MethodPost)
	router.HandleFunc("/api/users/registration/driver/verify/{username}", handler.DriverVerification).Methods(http.MethodPut)
	router.HandleFunc("/api/users/registration/passenger", handler.PassengerRegistration).Methods(http.MethodPost)

	router.HandleFunc("/api/users/search/admin", handler.SearchAdmin).Methods(http.MethodGet)
	router.HandleFunc("/api/users/search/driver", handler.SearchDriver).Methods(http.MethodGet)
	router.HandleFunc("/api/users/search/passenger", handler.SearchPassenger).Methods(http.MethodGet)

	router.HandleFunc("/api/users/admin/{username}", handler.GetAdmin).Methods(http.MethodGet)
	router.HandleFunc("/api/users/driver/{username}", handler.GetDriver).Methods(http.MethodGet)
	router.HandleFunc("/api/users/passenger/{username}", handler.GetPassenger).Methods(http.MethodGet)

	router.HandleFunc("/api/users/profile", handler.UpdateProfile).Methods(http.MethodPut)
	router.HandleFunc("/api/users/change-password", handler.ChangePassword).Methods(http.MethodPut)

	router.HandleFunc("/api/users/ban/driver/{username}", handler.BanDriver).Methods(http.MethodPut)
	router.HandleFunc("/api/users/ban/passenger/{username}", handler.BanPassenger).Methods(http.MethodPut)

	router.HandleFunc("/api/users/driver/{username}", handler.DeleteDriver).Methods(http.MethodDelete)
	router.HandleFunc("/api/users/passenger/{username}", handler.DeletePassenger).Methods(http.MethodDelete)

	router.HandleFunc("/api/users/role/{username}", handler.GetRoleForUsername).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8081", router))
}
