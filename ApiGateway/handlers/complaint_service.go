package handlers

import (
	"ApiGateway/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetAllComplaints(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "admin"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	SendReqAndReturnResponse(resWriter, req, http.MethodGet, _ComplaintServiceRoot)
}

func CreateComplaint(resWriter http.ResponseWriter, req *http.Request) {
	status, err := Authorize(req, "passenger")
	if err == nil {
		SendReqAndReturnResponse(resWriter, req, http.MethodPost, _ComplaintServiceRoot)
		return
	}

	status, err = Authorize(req, "driver")
	if err == nil {
		SendReqAndReturnResponse(resWriter, req, http.MethodPost, _ComplaintServiceRoot)
		return
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(status)
	json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
}

func DeleteComplaint(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "admin"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	params := mux.Vars(req)
	id := params["id"]

	SendReqAndReturnResponse(resWriter, req, http.MethodDelete, ComplaintServiceRoot+id)
}
