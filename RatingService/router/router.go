package router

import (
	"RatingService/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func MapRoutesAndServe(handler *handlers.RatingsHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/ratings", handler.CreateRating).Methods(http.MethodPost)
	router.HandleFunc("/api/ratings", handler.UpdateRating).Methods(http.MethodPut)
	router.HandleFunc("/api/ratings/{id}", handler.DeleteRating).Methods(http.MethodDelete)
	router.HandleFunc("/api/ratings/evaluated/{username}", handler.GetAllForEvaluated).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8084", router))
}
