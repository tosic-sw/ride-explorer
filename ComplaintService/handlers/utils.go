package handlers

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
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
