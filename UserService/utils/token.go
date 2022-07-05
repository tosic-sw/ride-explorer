package utils

import (
	"UserService/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var SECRET = []byte("sG9ncRwQJDrsr9jZ-rk4qMPT0t0ogbeLq")

func CreateToken(acc *models.UserAccount) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["expiresIn"] = time.Now().Add(time.Hour).Unix()
	claims["username"] = acc.Username
	claims["role"] = acc.Role

	return token.SignedString(SECRET)
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
