package handlers

import (
	"UserService/data"
	"UserService/models"
	"UserService/utils"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

type UsersHandler struct {
	repository *data.Repository
}

func NewUsersHandler(repository *data.Repository) *UsersHandler {
	return &UsersHandler{repository}
}

func (uh *UsersHandler) Login(resWriter http.ResponseWriter, req *http.Request) {
	var loginDTO models.LoginDTO
	json.NewDecoder(req.Body).Decode(&loginDTO)
	resWriter.Header().Set("Content-Type", "application/json")

	acc, err := uh.repository.FindOneLogin(loginDTO.Username)
	if err != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: err.Error()})
		return
	}

	if acc.Password != loginDTO.Password {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Wrong password"})
		return
	}

	tokenStr, err := utils.CreateToken(acc)
	if err != nil {
		fmt.Println(err.Error())
		resWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Unknown error happened at login"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.TokenState{Token: tokenStr})
}

func (uh *UsersHandler) Authenticate(resWriter http.ResponseWriter, req *http.Request) {
	bearer := req.Header["Authorization"]
	if bearer == nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenStr := strings.Split(bearer[0], " ")[1]
	token, err := utils.ParseTokenStr(tokenStr)
	claims := token.Claims.(jwt.MapClaims)

	if err != nil || !token.Valid {
		resWriter.WriteHeader(http.StatusUnauthorized)
	}

	// Check if banned or deleted
	var username = fmt.Sprintf("%v", claims["username"])
	_, err = uh.repository.FindOneLogin(username)
	if err != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
	}
}

func (uh *UsersHandler) AuthorizeAdmin(resWriter http.ResponseWriter, req *http.Request) {
	bearer := req.Header["Authorization"]
	if bearer == nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := strings.Split(bearer[0], " ")[1]
	token, err := utils.ParseTokenStr(tokenStr)
	claims := token.Claims.(jwt.MapClaims)

	if err != nil || !token.Valid || claims["role"] != models.ADMIN.String() {
		resWriter.WriteHeader(http.StatusUnauthorized)
	}

	// Check if banned or deleted
	var username = fmt.Sprintf("%v", claims["username"])
	_, err = uh.repository.FindOneLogin(username)
	if err != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
	}
}

func (uh *UsersHandler) AuthorizeDriver(resWriter http.ResponseWriter, req *http.Request) {
	bearer := req.Header["Authorization"]
	if bearer == nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := strings.Split(bearer[0], " ")[1]
	token, err := utils.ParseTokenStr(tokenStr)
	claims := token.Claims.(jwt.MapClaims)

	if err != nil || !token.Valid || claims["role"] != models.DRIVER.String() {
		resWriter.WriteHeader(http.StatusUnauthorized)
	}

	// Check if banned or deleted
	var username = fmt.Sprintf("%v", claims["username"])
	_, err = uh.repository.FindOneLogin(username)
	if err != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
	}
}

func (uh *UsersHandler) AuthorizePassenger(resWriter http.ResponseWriter, req *http.Request) {
	bearer := req.Header["Authorization"]
	if bearer == nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := strings.Split(bearer[0], " ")[1]
	token, err := utils.ParseTokenStr(tokenStr)
	claims := token.Claims.(jwt.MapClaims)

	if err != nil || !token.Valid || claims["role"] != models.PASSENGER.String() {
		resWriter.WriteHeader(http.StatusUnauthorized)
	}

	// Check if banned or deleted
	var username = fmt.Sprintf("%v", claims["username"])
	_, err = uh.repository.FindOneLogin(username)
	if err != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
	}
}

func (uh *UsersHandler) AdminRegistration(resWriter http.ResponseWriter, req *http.Request) {
	var regDTO models.RegistrationDTO
	resWriter.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(req.Body).Decode(&regDTO)
	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Invalid data sent"})
		return
	}

	if _, err = uh.repository.FindOneAcc(regDTO.Username); err == nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Username already taken"})
	}

	user := regDTO.ToAdmin()
	user, err = uh.repository.SaveAdmin(user)
	if err != nil {
		fmt.Println(err.Error())
		resWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Unknown error happened at admin registration"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Admin successfully registered"})
}

func (uh *UsersHandler) DriverRegistration(resWriter http.ResponseWriter, req *http.Request) {
	var regDTO models.DriverRegistrationDTO
	resWriter.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(req.Body).Decode(&regDTO)
	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Invalid data sent"})
		return
	}

	if _, err = uh.repository.FindOneAcc(regDTO.Username); err == nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Username already taken"})
		return
	}

	user := regDTO.ToDriver()
	user, err = uh.repository.SaveDriver(user)
	if err != nil {
		fmt.Println(err.Error())
		resWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Unknown error happened at driver registration"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "You have successfully sent registration request."})
}

func (uh *UsersHandler) DriverVerification(resWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	username := params["username"]

	resWriter.Header().Set("Content-Type", "application/json")

	account, err := uh.repository.FindOneAcc(username)
	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Driver does not exist in database"})
		return
	}

	if account.Verified {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Driver already verified"})
		return
	}

	err = uh.repository.VerifyDriver(username)
	if err != nil {
		fmt.Println(err.Error())
		resWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Unknown error happened while verifying driver"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Driver successfully verified"})
}

func (uh *UsersHandler) PassengerRegistration(resWriter http.ResponseWriter, req *http.Request) {
	var regDTO models.RegistrationDTO
	resWriter.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(req.Body).Decode(&regDTO)
	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Invalid data sent"})
		return
	}

	if _, err = uh.repository.FindOneAcc(regDTO.Username); err == nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Username already taken"})
		return
	}

	user := regDTO.ToPassenger()
	user, err = uh.repository.SavePassenger(user)
	if err != nil {
		fmt.Println(err.Error())
		resWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Unknown error happened at passenger registration"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "You have been successfully registered"})
}

func (uh *UsersHandler) SearchAdmin(resWriter http.ResponseWriter, req *http.Request) {
	search, offset, size := uh.parseSearchPageable(req)

	users, totalElements, _ := uh.repository.SearchAdmins(search, offset, size)
	userDTOs := []models.UserDTO{}

	for _, admin := range users {
		userDTOs = append(userDTOs, admin.ToDTO())
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.Header().Set("total-elements", strconv.FormatInt(totalElements, 10))
	json.NewEncoder(resWriter).Encode(userDTOs)
}

func (uh *UsersHandler) SearchDriver(resWriter http.ResponseWriter, req *http.Request) {
	search, offset, size := uh.parseSearchPageable(req)

	users, totalElements, _ := uh.repository.SearchDrivers(search, offset, size, true)
	userDTOs := []models.UserDTO{}

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
	userDTOs := []models.UserDTO{}

	for _, passenger := range users {
		userDTOs = append(userDTOs, passenger.ToDTO())
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.Header().Set("total-elements", strconv.FormatInt(totalElements, 10))
	json.NewEncoder(resWriter).Encode(userDTOs)
}

func (uh *UsersHandler) SearchUnverifiedDriver(resWriter http.ResponseWriter, req *http.Request) {
	search, offset, size := uh.parseSearchPageable(req)

	users, totalElements, _ := uh.repository.SearchDrivers(search, offset, size, false)
	userDTOs := []models.UserDTO{}

	for _, driver := range users {
		userDTOs = append(userDTOs, driver.ToDTO())
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.Header().Set("total-elements", strconv.FormatInt(totalElements, 10))
	json.NewEncoder(resWriter).Encode(userDTOs)
}

func (uh *UsersHandler) GetAdmin(resWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	username := params["username"]

	resWriter.Header().Set("Content-Type", "application/json")

	admin, err := uh.repository.FindOneAdmin(username)
	if err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(admin.ToDTO())
}

func (uh *UsersHandler) GetDriver(resWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	username := params["username"]

	resWriter.Header().Set("Content-Type", "application/json")

	driver, err := uh.repository.FindOneDriverWithCar(username)

	if err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Driver not found"})
		return
	}

	json.NewEncoder(resWriter).Encode(driver.ToDriverDTO())
}

func (uh *UsersHandler) GetUnverifiedDriver(resWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	username := params["username"]

	resWriter.Header().Set("Content-Type", "application/json")

	driver, err := uh.repository.FindOneUnverifiedDriverWithCar(username)
	if err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Unverified driver not found for given username"})
		return
	}

	json.NewEncoder(resWriter).Encode(driver.ToDriverDTO())
}

func (uh *UsersHandler) GetPassenger(resWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	username := params["username"]

	resWriter.Header().Set("Content-Type", "application/json")

	passenger, err := uh.repository.FindOnePassenger(username)
	if err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(passenger.ToDTO())
}

func (uh *UsersHandler) UpdateProfile(resWriter http.ResponseWriter, req *http.Request) {
	var userDTO models.UserForUpdateDTO
	resWriter.Header().Set("Content-Type", "application/json")

	var err = json.NewDecoder(req.Body).Decode(&userDTO)
	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Invalid data sent"})
		return
	}

	// Dobavi claims
	bearer := req.Header["Authorization"]
	if bearer == nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Unauthorized"})
		return
	}

	tokenStr := strings.Split(bearer[0], " ")[1]
	token, err := utils.ParseTokenStr(tokenStr)
	claims := token.Claims.(jwt.MapClaims)

	var username = fmt.Sprintf("%v", claims["username"])

	if claims["role"] == models.ADMIN.String() {
		_, err = uh.repository.FindOneAdmin(username)
	} else if claims["role"] == models.DRIVER.String() {
		_, err = uh.repository.FindOneDriver(username)
	} else if claims["role"] == models.PASSENGER.String() {
		_, err = uh.repository.FindOnePassenger(username)
	}

	if err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: err.Error()})
		return
	}

	if claims["role"] == models.ADMIN.String() {
		_, err = uh.repository.UpdateAdmin(&userDTO, username)
	} else if claims["role"] == models.DRIVER.String() {
		_, err = uh.repository.UpdateDriver(&userDTO, username)
	} else if claims["role"] == models.PASSENGER.String() {
		_, err = uh.repository.UpdatePassenger(&userDTO, username)
	}

	if err != nil {
		fmt.Println(err.Error())
		resWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Profile successfully updated"})
}

func (uh *UsersHandler) ChangePassword(resWriter http.ResponseWriter, req *http.Request) {
	var userDTO models.UserPasswordChangeDTO
	resWriter.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(req.Body).Decode(&userDTO)
	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Invalid data sent"})
		return
	}

	bearer := req.Header["Authorization"]
	if bearer == nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Unauthorized"})
		return
	}

	tokenStr := strings.Split(bearer[0], " ")[1]
	token, err := utils.ParseTokenStr(tokenStr)
	claims := token.Claims.(jwt.MapClaims)

	var username = fmt.Sprintf("%v", claims["username"])

	_, err = uh.repository.FindOneAcc(username)
	if err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: err.Error()})
		return
	}

	_, err = uh.repository.ChangePassword(username, userDTO.Password)
	if err != nil {
		fmt.Println(err.Error())
		resWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Password successfully changed"})
}

func (uh *UsersHandler) BanDriver(resWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	username := params["username"]

	resWriter.Header().Set("Content-Type", "application/json")

	_, _, err := uh.repository.BanDriver(username)
	if err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Driver successfully banned"})
}

func (uh *UsersHandler) BanPassenger(resWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	username := params["username"]

	resWriter.Header().Set("Content-Type", "application/json")

	_, _, err := uh.repository.BanPassenger(username)
	if err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Passenger successfully banned"})
}

func (uh *UsersHandler) DeleteDriver(resWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	username := params["username"]

	resWriter.Header().Set("Content-Type", "application/json")

	err := uh.repository.DeleteDriver(username)
	if err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Driver to delete not found"})
		return
	}
	_ = uh.repository.DeleteUserAccount(username)

	json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Driver successfully deleted"})
}

func (uh *UsersHandler) DeletePassenger(resWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	username := params["username"]

	resWriter.Header().Set("Content-Type", "application/json")

	err := uh.repository.DeletePassenger(username)
	if err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Passenger to delete not found"})
		return
	}
	_ = uh.repository.DeleteUserAccount(username)

	json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "Passenger successfully deleted"})
}

func (uh *UsersHandler) GetRoleForUsername(resWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	username := params["username"]

	resWriter.Header().Set("Content-Type", "application/json")

	acc, err := uh.repository.FindOneAccRole(username)
	if err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.MessageResponse{Message: "User with given username not found"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.RoleDTO{Role: acc.Role.String()})
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
