package handlers

import (
	"ApiGateway/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetDrive(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authenticate(req); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	params := mux.Vars(req)
	id := params["id"]

	SendReqAndReturnResponse(resWriter, req, http.MethodGet, DriveServiceRoot+id)
}

func GetUnfinishedDriveForDriver(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authenticate(req); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	params := mux.Vars(req)
	id := params["id"]
	username := params["username"]

	SendReqAndReturnResponse(resWriter, req, http.MethodGet, DriveServiceRoot+Unfinished+id+"/"+username)
}

func CreateDrive(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "driver"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	SendReqAndReturnResponse(resWriter, req, http.MethodPost, _DriveServiceRoot)
}

func UpdateDrive(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "driver"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	SendReqAndReturnResponse(resWriter, req, http.MethodPut, _DriveServiceRoot)
}

func FinishDriveOfDriver(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "driver"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	params := mux.Vars(req)
	username := params["username"]
	id := params["id"]

	SendReqAndReturnResponse(resWriter, req, http.MethodPut, DriveServiceRoot+Driver+username+Slash+Finish+id)
}

func AdjustPlaces(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "passenger"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	SendReqAndReturnResponse(resWriter, req, http.MethodPut, DriveServiceRoot+_Reserve)
}

func DeleteDriveOfDriver(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "driver"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	params := mux.Vars(req)
	username := params["username"]
	id := params["id"]

	SendReqAndReturnResponse(resWriter, req, http.MethodDelete, DriveServiceRoot+Driver+username+Slash+id)
}

func SearchDrives(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "passenger"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	SendReqAndReturnResponse(resWriter, req, http.MethodPost, DriveServiceRoot+_Search)
}

func FinishedDrivesOfDriver(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "driver"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	params := mux.Vars(req)
	username := params["username"]

	SendReqAndReturnResponse(resWriter, req, http.MethodPost, DriveServiceRoot+Driver+Finished+username)
}

func UnfinishedDrivesOfDriver(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "driver"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	params := mux.Vars(req)
	username := params["username"]

	SendReqAndReturnResponse(resWriter, req, http.MethodPost, DriveServiceRoot+Driver+Unfinished+username)
}
