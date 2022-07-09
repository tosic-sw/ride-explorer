package router

import (
	"ComplaintService/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func MapRoutesAndServe(handler *handlers.ComplaintsHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/complaints", handler.GetAllComplaints).Methods(http.MethodGet)
	router.HandleFunc("/api/complaints", handler.CreateComplaint).Methods(http.MethodPost)
	router.HandleFunc("/api/complaints/{id:[0-9]+}", handler.DeleteComplaint).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8083", router))
}
