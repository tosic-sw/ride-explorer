package handlers

import (
	"ApiGateway/models"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

func Authorize(req *http.Request, role string) (int, error) {
	bearer := req.Header["Authorization"]
	if bearer == nil {
		return http.StatusBadRequest, errors.New("no token bearer")
	}
	tokenStr := strings.Split(bearer[0], " ")[1]

	timeout := time.Duration(5 * time.Second)
	client := http.Client{Timeout: timeout}

	endpoint := UserServiceRoot + AuthorizeApi + role
	request, _ := http.NewRequest(http.MethodGet, endpoint, bytes.NewBufferString(""))
	request.Header.Set("Authorization", "Bearer "+tokenStr)

	response, err := client.Do(request)
	if err != nil {
		return http.StatusGatewayTimeout, errors.New("error sending request")
	}

	if response.StatusCode != 200 {
		return http.StatusUnauthorized, errors.New("unauthorized")
	}

	return http.StatusOK, nil
}

func Authenticate(req *http.Request) (int, error) {
	bearer := req.Header["Authorization"]
	if bearer == nil {
		return http.StatusBadRequest, errors.New("no token bearer")
	}
	tokenStr := strings.Split(bearer[0], " ")[1]

	timeout := time.Duration(5 * time.Second)
	client := http.Client{Timeout: timeout}

	endpoint := UserServiceRoot + _AuthenticateApi
	request, _ := http.NewRequest(http.MethodGet, endpoint, bytes.NewBufferString(""))
	request.Header.Set("Authorization", "Bearer "+tokenStr)

	response, err := client.Do(request)
	if err != nil {
		return http.StatusGatewayTimeout, errors.New("error sending request")
	}

	if response.StatusCode != 200 {
		return http.StatusUnauthorized, errors.New("not authenticated")
	}

	return http.StatusOK, nil
}

func SendReqAndReturnResponse(resWriter http.ResponseWriter, req *http.Request, method string, endpoint string) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{Timeout: timeout}

	request, _ := http.NewRequest(method, endpoint, req.Body)
	if bearer := req.Header["Authorization"]; bearer != nil {
		request.Header.Set("Authorization", bearer[0])
	}

	response, err := client.Do(request)
	if err != nil {
		resWriter.Header().Set("Content-Type", response.Header.Get("application/json"))
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		json.NewEncoder(resWriter).Encode(models.ErrorResponse{Message: "Unknown error happened"})
		return
	}

	AdjustAllHeaders(&resWriter, response)
	io.Copy(resWriter, response.Body)
	response.Body.Close()
}

func AdjustAllHeaders(resWriter *http.ResponseWriter, response *http.Response) {
	for name, values := range response.Header {
		for _, value := range values {
			(*resWriter).Header().Set(name, value)
		}
	}
	(*resWriter).WriteHeader(response.StatusCode)
}

func GetSearchPageableFromRequest(req *http.Request) string {
	search := req.URL.Query().Get("search")
	size := req.URL.Query().Get("size")
	page := req.URL.Query().Get("page")

	return composeSearchPageable(search, size, page)
}

func composeSearchPageable(search string, size string, page string) string {
	return QMark + pSearch + search + Amp + pSize + size + Amp + pPage + page
}

func GetPageableFromRequest(req *http.Request) string {
	size := req.URL.Query().Get("size")
	page := req.URL.Query().Get("page")

	return composePageable(size, page)
}

func composePageable(size string, page string) string {
	return QMark + pSize + size + Amp + pPage + page
}
