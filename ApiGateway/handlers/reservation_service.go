package handlers

import (
	"ApiGateway/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetReservation(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "passenger"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	params := mux.Vars(req)
	id := params["id"]

	SendReqAndReturnResponse(resWriter, req, http.MethodGet, ReservationServiceRoot+id)
}

func CreateReservation(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "passenger"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	SendReqAndReturnResponse(resWriter, req, http.MethodPost, _ReservationServiceRoot)
}

func DeleteReservation(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "passenger"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	params := mux.Vars(req)
	id := params["id"]

	SendReqAndReturnResponse(resWriter, req, http.MethodDelete, ReservationServiceRoot+id)
}

func VerifyReservation(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "driver"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	params := mux.Vars(req)
	id := params["id"]

	SendReqAndReturnResponse(resWriter, req, http.MethodPut, ReservationServiceRoot+Verify+id)
}

func GetAllByUserVerified(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "passenger"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	SendReqAndReturnResponse(resWriter, req, http.MethodGet, ReservationServiceRoot+User+_Verified)
}

func GetAllByUserUnverified(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "passenger"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	SendReqAndReturnResponse(resWriter, req, http.MethodGet, ReservationServiceRoot+User+_Unverified)
}
