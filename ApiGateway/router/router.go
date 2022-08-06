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

	router.HandleFunc("/api/drives/{id}", handlers.GetDrive).Methods(http.MethodGet)
	router.HandleFunc("/api/drives", handlers.CreateDrive).Methods(http.MethodPost)
	router.HandleFunc("/api/drives", handlers.UpdateDrive).Methods(http.MethodPut)
	router.HandleFunc("/api/drives/driver/{username}/finish/{id}", handlers.FinishDriveOfDriver).Methods(http.MethodPut)
	router.HandleFunc("/api/drives/adjust-places", handlers.AdjustPlaces).Methods(http.MethodPut)
	router.HandleFunc("/api/drives/driver/{username}/{id}", handlers.DeleteDriveOfDriver).Methods(http.MethodDelete)
	router.HandleFunc("/api/drives/search", handlers.SearchDrives).Methods(http.MethodPost)
	router.HandleFunc("/api/drives/driver/finished/{username}", handlers.FinishedDrivesOfDriver).Methods(http.MethodGet)
	router.HandleFunc("/api/drives/driver/unfinished/{username}", handlers.UnfinishedDrivesOfDriver).Methods(http.MethodGet)

	router.HandleFunc("/api/reservations/{id:[0-9]+}", handlers.GetReservation).Methods(http.MethodGet)
	router.HandleFunc("/api/reservations", handlers.CreateReservation).Methods(http.MethodPost)
	router.HandleFunc("/api/reservations/{id:[0-9]+}", handlers.DeleteReservation).Methods(http.MethodDelete)
	router.HandleFunc("/api/reservations/verify/{id:[0-9]+}", handlers.VerifyReservation).Methods(http.MethodPut)
	router.HandleFunc("/api/reservations/user/verified", handlers.GetAllByUserVerified).Methods(http.MethodGet)
	router.HandleFunc("/api/reservations/user/unverified", handlers.GetAllByUserUnverified).Methods(http.MethodGet)

	router.HandleFunc("/api/complaints", handlers.GetAllComplaints).Methods(http.MethodGet)
	router.HandleFunc("/api/complaints", handlers.CreateComplaint).Methods(http.MethodPost)
	router.HandleFunc("/api/complaints/{id:[0-9]+}", handlers.DeleteComplaint).Methods(http.MethodDelete)

	router.HandleFunc("/api/ratings", handlers.CreateRating).Methods(http.MethodPost)
	router.HandleFunc("/api/ratings/{id:[0-9]+}", handlers.UpdateRating).Methods(http.MethodPut)
	router.HandleFunc("/api/ratings/{id:[0-9]+}", handlers.DeleteRating).Methods(http.MethodDelete)
	router.HandleFunc("/api/ratings/evaluated/{username}", handlers.GetAllForEvaluated).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}
