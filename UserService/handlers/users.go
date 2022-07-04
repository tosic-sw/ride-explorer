package handlers

import (
	"UserService/data"
	"net/http"
)

type UsersHandler struct {
	repository *data.Repository
}

func NewUsersHandler(repository *data.Repository) *UsersHandler {
	return &UsersHandler{repository}
}

func (uh *UsersHandler) Login(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) DriverRegistration(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) PassengerRegistration(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) AdminRegistration(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) SearchAdmin(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) SearchDriver(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) SearchPassenger(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) GetAdmin(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) GetDriver(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) GetPassenger(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) UpdateUserData(resWriter http.ResponseWriter, req *http.Request) {

}
