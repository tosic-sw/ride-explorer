package handlers

import (
	"UserService/data"
	"UserService/models"
	"encoding/json"
	"net/http"
	"strconv"
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
	search, offset, size := uh.parseSearchPageable(req)

	users, totalElements, _ := uh.repository.SearchAdmins(search, offset, size)
	var userDTOs []models.UserDTO

	for _, admin := range users {
		userDTOs = append(userDTOs, admin.ToDTO())
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.Header().Set("total-elements", strconv.FormatInt(totalElements, 10))
	json.NewEncoder(resWriter).Encode(userDTOs)
}

func (uh *UsersHandler) SearchDriver(resWriter http.ResponseWriter, req *http.Request) {
	search, offset, size := uh.parseSearchPageable(req)

	users, totalElements, _ := uh.repository.SearchDrivers(search, offset, size)
	var userDTOs []models.UserDTO

	for _, driver := range users {
		userDTOs = append(userDTOs, driver.ToDTO())
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.Header().Set("total-elements", strconv.FormatInt(totalElements, 10))
	json.NewEncoder(resWriter).Encode(userDTOs)
}

func (uh *UsersHandler) SearchPassenger(resWriter http.ResponseWriter, req *http.Request) {
	search, offset, size := uh.parseSearchPageable(req)

	users, totalElements, _ := uh.repository.SearchPassengers(search, offset, size)
	var userDTOs []models.UserDTO

	for _, passenger := range users {
		userDTOs = append(userDTOs, passenger.ToDTO())
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.Header().Set("total-elements", strconv.FormatInt(totalElements, 10))
	json.NewEncoder(resWriter).Encode(userDTOs)
}

func (uh *UsersHandler) GetAdmin(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) GetDriver(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) GetPassenger(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) UpdateAdmin(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) UpdateDriver(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) UpdatePassenger(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) BanDriver(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) BanPassenger(resWriter http.ResponseWriter, req *http.Request) {

}

func (uh *UsersHandler) parseSearchPageable(req *http.Request) (string, int, int) {
	q := req.URL.Query()
	search := q.Get("search")

	pageStr := q.Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	if page == 0 {
		page = 1
	}

	size, err := strconv.Atoi(q.Get("size"))
	if err != nil {
		size = 5
	}
	switch {
	case size > 100:
		size = 100
	case size <= 0:
		size = 10
	}

	offset := (page - 1) * size
	return search, offset, size
}
