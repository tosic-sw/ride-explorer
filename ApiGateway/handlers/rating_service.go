package handlers

import (
	"ApiGateway/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateRating(resWriter http.ResponseWriter, req *http.Request) {
	status, err := Authorize(req, "passenger")
	if err == nil {
		SendReqAndReturnResponse(resWriter, req, http.MethodPost, _RatingServiceRoot)
		return
	}

	status, err = Authorize(req, "driver")
	if err == nil {
		SendReqAndReturnResponse(resWriter, req, http.MethodPost, _RatingServiceRoot)
		return
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(status)
	json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
}

func UpdateRating(resWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	status, err := Authorize(req, "passenger")
	if err == nil {
		SendReqAndReturnResponse(resWriter, req, http.MethodPut, RatingServiceRoot+id)
		return
	}

	status, err = Authorize(req, "driver")
	if err == nil {
		SendReqAndReturnResponse(resWriter, req, http.MethodPut, RatingServiceRoot+id)
		return
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(status)
	json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
}

func DeleteRating(resWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	status, err := Authorize(req, "passenger")
	if err == nil {
		SendReqAndReturnResponse(resWriter, req, http.MethodDelete, RatingServiceRoot+id)
		return
	}

	status, err = Authorize(req, "driver")
	if err == nil {
		SendReqAndReturnResponse(resWriter, req, http.MethodDelete, RatingServiceRoot+id)
		return
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(status)
	json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
}

func GetAllForEvaluated(resWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	username := params["username"]
	pageable := GetPageableFromRequest(req)

	status, err := Authenticate(req)
	if err == nil {
		SendReqAndReturnResponse(resWriter, req, http.MethodGet, RatingServiceRoot+Evaluated+username+pageable)
		return
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(status)
	json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
}
