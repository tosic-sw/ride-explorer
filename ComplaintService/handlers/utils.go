package handlers

import (
	"ComplaintService/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"strings"
)

var SECRET = []byte("sG9ncRwQJDrsr9jZ-rk4qMPT0t0ogbeLq")

func AdjustResponseHeaderJson(resWriter *http.ResponseWriter) {
	(*resWriter).Header().Set("Content-Type", "application/json")
}

func AdjustResponseHeaderText(resWriter *http.ResponseWriter) {
	(*resWriter).Header().Set("Content-Type", "application/json")
}

func ParseTokenStr(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("error while parsing jwt str")
		}
		return SECRET, nil
	})

	return token, err
}

func GetUsernameFromRequest(req *http.Request) (string, error) {
	bearer := req.Header["Authorization"]
	if bearer == nil {
		return "", errors.New("unauthorized")
	}
	tokenStr := strings.Split(bearer[0], " ")[1]
	token, _ := ParseTokenStr(tokenStr) // Token is always valid because of authorization
	claims := token.Claims.(jwt.MapClaims)

	return fmt.Sprintf("%v", claims["username"]), nil
}

func GetRoleFromRequest(req *http.Request) (string, error) {
	bearer := req.Header["Authorization"]
	if bearer == nil {
		return "", errors.New("unauthorized")
	}
	tokenStr := strings.Split(bearer[0], " ")[1]
	token, _ := ParseTokenStr(tokenStr) // Token is always valid because of authorization
	claims := token.Claims.(jwt.MapClaims)

	return fmt.Sprintf("%v", claims["role"]), nil
}

func GetRoleOfUser(username string) (string, error) {
	endpoint := "http://localhost:8081/api/users/role" + "/" + username
	resp, err := http.Get(endpoint)

	if err != nil {
		return "", errors.New("internal server error")
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("invalid user for complaint")
	}

	var roleDTO models.RoleDTO
	err = json.NewDecoder(resp.Body).Decode(&roleDTO)
	if err != nil {
		return "", errors.New("internal server error")
	}

	return roleDTO.Role, nil
}

func ExistsFinishedDrive(id uint) error {
	endpoint := "http://localhost:8000/api/drives/" + strconv.Itoa(int(id))
	resp, err := http.Get(endpoint)
	if err != nil {
		return errors.New("internal server error")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("unfinished or invalid drive")
	}

	return nil
}

func ExistsFinishedDriveDriver(driveId uint, driver string) error {
	endpoint := "http://localhost:8000/api/drives/" + strconv.Itoa(int(driveId)) + "/" + driver
	resp, err := http.Get(endpoint)
	if err != nil {
		return errors.New("internal server error")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("unfinished drive or invalid driver")
	}

	return nil
}

func ExistsVerifiedReservation(driveId uint, passenger string) error {
	endpoint := "http://localhost:8082/api/reservations/is-verified/" + passenger + "/" + strconv.Itoa(int(driveId))
	resp, err := http.Get(endpoint)
	if err != nil {
		return errors.New("internal server error")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("reservation does not match with drive, or is not verified")
	}

	return nil
}
