package handlers

import (
	"RatingService/data"
	"RatingService/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type RatingsHandler struct {
	repository *data.Repository
}

func NewRatingsHandler(repository *data.Repository) *RatingsHandler {
	return &RatingsHandler{repository}
}

func (ch *RatingsHandler) CreateRating(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)
	username, err := GetUsernameFromRequest(req)
	if err != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Unauthorized"})
		return
	}

	var ratingDTO models.RatingDTO
	err = json.NewDecoder(req.Body).Decode(&ratingDTO)
	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Invalid data sent"})
		return
	}

	_, err = ch.repository.FindOneComplex(username, ratingDTO.Evaluated, ratingDTO.DriveId)
	if err == nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Rating already exists for drive"})
		return
	}

	role, err := GetRoleFromRequest(req)
	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: err.Error()})
		return
	}

	rating := ratingDTO.ToRating(username)
	if role == "DRIVER" {
		ch.ProceedSavingAndReturnResponseDriver(&resWriter, &rating)
	} else if role == "PASSENGER" {
		ch.ProceedSavingAndReturnResponsePassenger(&resWriter, &rating)
	}
}

func (ch *RatingsHandler) ProceedSavingAndReturnResponseDriver(resWriter *http.ResponseWriter, rating *models.Rating) {
	if err := ExistsFinishedDriveDriver(rating.DriveId, rating.Evaluator); err != nil {
		json.NewEncoder(*resWriter).Encode(models.Response{Message: err.Error()})
		return
	}

	if err := ExistsVerifiedReservation(rating.DriveId, rating.Evaluated); err != nil {
		json.NewEncoder(*resWriter).Encode(models.Response{Message: err.Error()})
		return
	}

	_, err := ch.repository.SaveRating(rating)
	if err != nil {
		fmt.Println(err.Error())
		(*resWriter).WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(*resWriter).Encode(models.Response{Message: "Unknown error happened while saving rating"})
	}

	json.NewEncoder(*resWriter).Encode(models.Response{Message: "Rating successfully created"})
}

func (ch *RatingsHandler) ProceedSavingAndReturnResponsePassenger(resWriter *http.ResponseWriter, rating *models.Rating) {
	evaluatedRole, err := GetRoleOfUser(rating.Evaluated)
	if err != nil {
		(*resWriter).WriteHeader(http.StatusBadRequest)
		json.NewEncoder(*resWriter).Encode(models.Response{Message: err.Error()})
		return
	}

	if evaluatedRole == "DRIVER" {
		if err := ExistsFinishedDriveDriver(rating.DriveId, rating.Evaluated); err != nil {
			(*resWriter).WriteHeader(http.StatusBadRequest)
			json.NewEncoder(*resWriter).Encode(models.Response{Message: err.Error()})
			return
		}
		if err := ExistsVerifiedReservation(rating.DriveId, rating.Evaluator); err != nil {
			(*resWriter).WriteHeader(http.StatusBadRequest)
			json.NewEncoder(*resWriter).Encode(models.Response{Message: err.Error()})
			return
		}
	} else if evaluatedRole == "PASSENGER" {
		if err := ExistsFinishedDrive(rating.DriveId); err != nil {
			(*resWriter).WriteHeader(http.StatusBadRequest)
			json.NewEncoder(*resWriter).Encode(models.Response{Message: err.Error()})
			return
		}
		if err := ExistsVerifiedReservation(rating.DriveId, rating.Evaluator); err != nil {
			(*resWriter).WriteHeader(http.StatusBadRequest)
			json.NewEncoder(*resWriter).Encode(models.Response{Message: err.Error()})
			return
		}
		if err := ExistsVerifiedReservation(rating.DriveId, rating.Evaluated); err != nil {
			(*resWriter).WriteHeader(http.StatusBadRequest)
			json.NewEncoder(*resWriter).Encode(models.Response{Message: err.Error()})
			return
		}
	}

	_, err = ch.repository.SaveRating(rating)
	if err != nil {
		fmt.Println(err.Error())
		(*resWriter).WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(*resWriter).Encode(models.Response{Message: "Unknown error happened while saving rating"})
		return
	}

	json.NewEncoder(*resWriter).Encode(models.Response{Message: "Rating successfully created"})
}

func (ch *RatingsHandler) DeleteRating(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	username, err := GetUsernameFromRequest(req)

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: err.Error()})
		return
	}

	if _, err := ch.repository.FindOneByEvaluator(uint(idInt), username); err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Invalid rating id sent for given evaluator"})
		return
	}

	err = ch.repository.DeleteRating(uint(idInt))
	if err != nil {
		fmt.Println(err.Error())
		resWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Unknown error happened while deleting rating"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.Response{Message: "Rating successfully deleted"})
}

func (ch *RatingsHandler) UpdateRating(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, err := strconv.ParseInt(idStr, 10, 64)

	username, err := GetUsernameFromRequest(req)
	var ratingDTO models.RatingDTO
	err = json.NewDecoder(req.Body).Decode(&ratingDTO)
	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Invalid data sent"})
		return
	}

	if _, err := ch.repository.FindOneComplex(username, ratingDTO.Evaluated, ratingDTO.DriveId); err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Rating with given params does not exist"})
		return
	}

	_, err = ch.repository.UpdateRating(uint(idInt), username, ratingDTO.Positive, ratingDTO.Text)
	if err != nil {
		fmt.Println(err.Error())
		resWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Unknown error happened while updating rating"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.Response{Message: "Rating successfully updated"})
}

func (ch *RatingsHandler) GetAllForEvaluated(resWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	evaluated := params["username"]

	offset, size := ch.parseSearchPageable(req)

	ratings, totalElements, _ := ch.repository.FindAllForEvaluated(evaluated, offset, size)
	var resDTOs []models.ViewRatingDTO

	for _, rating := range ratings {
		resDTOs = append(resDTOs, rating.ToDTO())
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.Header().Set("total-elements", strconv.FormatInt(totalElements, 10))
	json.NewEncoder(resWriter).Encode(resDTOs)
}

func (ch *RatingsHandler) parseSearchPageable(req *http.Request) (int, int) {
	q := req.URL.Query()

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
	return offset, size
}
