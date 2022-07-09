package handlers

import (
	"ReservationService/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func AdjustResponseHeaderJson(resWriter *http.ResponseWriter) {
	(*resWriter).Header().Set("Content-Type", "application/json")
}

func AdjustResponseHeaderText(resWriter *http.ResponseWriter) {
	(*resWriter).Header().Set("Content-Type", "application/json")
}

var SECRET = []byte("sG9ncRwQJDrsr9jZ-rk4qMPT0t0ogbeLq")

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

func UpdateDrivePlaces(driveId int32, places int32) error {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{Timeout: timeout}

	dto := models.ReserveDTO{
		Id:     driveId,
		Places: places,
	}
	jsonData, err := json.Marshal(dto)

	endpoint := "http://localhost:8000/api/drives/adjust-places"
	request, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	resp, err := client.Do(request)
	if resp.StatusCode != http.StatusOK {
		return errors.New("invalid request, could not update places")
	}

	return nil
}

func VerifyDriveReservation(driveId int32, driver_username string) error {
	driveIdStr := strconv.Itoa(int(driveId))

	endpoint := "http://localhost:8000/api/drives/unfinished" + "/" + driveIdStr + "/" + driver_username
	resp, err := http.Get(endpoint)

	if resp.StatusCode != http.StatusOK || err != nil {
		return errors.New("invalid request, could not update places")
	}

	return nil
}
