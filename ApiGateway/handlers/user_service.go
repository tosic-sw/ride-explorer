package handlers

import (
	"ApiGateway/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"time"
)

func Login(resWriter http.ResponseWriter, req *http.Request) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{Timeout: timeout}

	request, _ := http.NewRequest(http.MethodPost, UserServiceRoot+LoginApi, req.Body)

	response, err := client.Do(request)
	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	AdjustAllHeaders(&resWriter, response)
	io.Copy(resWriter, response.Body)
	response.Body.Close()
}

func RegistrationAdmin(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "admin"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	SendReqAndReturnResponse(resWriter, req, http.MethodPost, UserServiceRoot+Registration+_Admin)
}

func RegistrationDriver(resWriter http.ResponseWriter, req *http.Request) {
	SendReqAndReturnResponse(resWriter, req, http.MethodPost, UserServiceRoot+Registration+_Driver)
}

func DriverVerification(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "admin"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}
	params := mux.Vars(req)
	username := params["username"]

	SendReqAndReturnResponse(resWriter, req, http.MethodPut, UserServiceRoot+Registration+Driver+Verify+username)
}

func RegistrationPassenger(resWriter http.ResponseWriter, req *http.Request) {
	SendReqAndReturnResponse(resWriter, req, http.MethodPost, UserServiceRoot+Registration+_Driver)
}

func SearchAdmin(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "admin"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	pageable := GetPageableFromRequest(req)
	SendReqAndReturnResponse(resWriter, req, http.MethodGet, UserServiceRoot+Search+_Admin+pageable)
}

func SearchDriver(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "admin"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	pageable := GetPageableFromRequest(req)
	SendReqAndReturnResponse(resWriter, req, http.MethodGet, UserServiceRoot+Search+_Driver+pageable)
}

func SearchPassenger(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "admin"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	pageable := GetPageableFromRequest(req)
	SendReqAndReturnResponse(resWriter, req, http.MethodGet, UserServiceRoot+Search+_Passenger+pageable)
}

func GetAdmin(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "admin"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}
	params := mux.Vars(req)
	username := params["username"]

	SendReqAndReturnResponse(resWriter, req, http.MethodGet, UserServiceRoot+Admin+username)
}

func GetDriver(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authenticate(req); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}
	params := mux.Vars(req)
	username := params["username"]

	SendReqAndReturnResponse(resWriter, req, http.MethodGet, UserServiceRoot+Driver+username)
}

func GetPassenger(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authenticate(req); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}
	params := mux.Vars(req)
	username := params["username"]

	SendReqAndReturnResponse(resWriter, req, http.MethodGet, UserServiceRoot+Passenger+username)
}

func UpdateProfile(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authenticate(req); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	SendReqAndReturnResponse(resWriter, req, http.MethodPut, UserServiceRoot+_Profile)
}

func ChangePassword(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authenticate(req); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}

	SendReqAndReturnResponse(resWriter, req, http.MethodPut, UserServiceRoot+_ChangePassword)
}

func BanDriver(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "admin"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}
	params := mux.Vars(req)
	username := params["username"]

	SendReqAndReturnResponse(resWriter, req, http.MethodPut, UserServiceRoot+Ban+Driver+username)
}

func BanPassenger(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "admin"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}
	params := mux.Vars(req)
	username := params["username"]

	SendReqAndReturnResponse(resWriter, req, http.MethodPut, UserServiceRoot+Ban+Passenger+username)
}

func DeleteDriver(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "admin"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}
	params := mux.Vars(req)
	username := params["username"]

	SendReqAndReturnResponse(resWriter, req, http.MethodDelete, UserServiceRoot+Driver+username)
}

func DeletePassenger(resWriter http.ResponseWriter, req *http.Request) {
	if status, err := Authorize(req, "admin"); err != nil {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(status)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}
	params := mux.Vars(req)
	username := params["username"]

	SendReqAndReturnResponse(resWriter, req, http.MethodDelete, UserServiceRoot+Passenger+username)
}
