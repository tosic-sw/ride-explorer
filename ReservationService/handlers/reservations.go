package handlers

import (
	"ReservationService/data"
	"ReservationService/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ReservationsHandler struct {
	repository  *data.Repository
	mailHandler *MailHandler
}

func NewReservationsHandler(repository *data.Repository, mh *MailHandler) *ReservationsHandler {
	return &ReservationsHandler{
		repository,
		mh,
	}
}

func (uh *ReservationsHandler) GetReservation(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, err := strconv.ParseInt(idStr, 10, 64)

	username, err := GetUsernameFromRequest(req)
	if err != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Unauthorized"})
		return
	}

	reservation, err := uh.repository.FindOneByUser(uint(idInt), username)
	if err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Invalid reservation id or not right permission to access it"})
		return
	}
	resDTO := reservation.ToDTO()

	json.NewEncoder(resWriter).Encode(resDTO)
}

func (uh *ReservationsHandler) CreateReservation(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)
	username, err := GetUsernameFromRequest(req)
	if err != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Unauthorized"})
		return
	}

	var resDTO models.CreateReservationDTO
	err = json.NewDecoder(req.Body).Decode(&resDTO)
	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Invalid data sent"})
		return
	}

	err = VerifyDriveReservation(resDTO.DriveId, resDTO.DriverUsername)
	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Invalid drive data"})
		return
	}

	res := resDTO.ToReservation(username)
	_, err = uh.repository.SaveReservation(res)
	if err != nil {
		fmt.Println(err.Error())
		resWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Unknown error happened while creating reservations"})
	}

	json.NewEncoder(resWriter).Encode(models.Response{Message: "Reservation successfully created"})
}

func (uh *ReservationsHandler) DeleteReservation(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, err := strconv.ParseInt(idStr, 10, 64)

	username, err := GetUsernameFromRequest(req)
	if err != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Unauthorized"})
		return
	}

	res, err := uh.repository.FindOneByUser(uint(idInt), username)
	if err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Invalid reservation id or not right permission to delete it"})
		return
	}

	err = UpdateDrivePlaces(res.DriveId, -1)
	if err != nil {
		resWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Reservation deletion failed."})
		return
	}

	err = uh.repository.DeleteReservation(uint(idInt), username)
	if err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Invalid reservation id or not right permission to delete it"})
		return
	}

	mailMessage := models.Mail{
		To:      res.GetDriverEmail(),
		Subject: "Ride data changed",
		Body:    ComposeReservationDeletedEmail(int(res.DriveId), res.DriverUsername, res.PassengerUsername),
	}
	uh.mailHandler.SendMail(mailMessage)

	json.NewEncoder(resWriter).Encode(models.Response{Message: "Reservation successfully deleted"})
}

func (uh *ReservationsHandler) VerifyReservation(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, err := strconv.ParseInt(idStr, 10, 64)

	username, err := GetUsernameFromRequest(req)
	if err != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Unauthorized"})
		return
	}

	res, err := uh.repository.FindOne(uint(idInt))
	if err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Invalid reservation id or not right permission to verify it"})
		return
	}

	err = UpdateDrivePlaces(res.DriveId, 1)
	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Reservation verification failed. Check free places again"})
		return
	}

	_, err = uh.repository.VerifyReservation(uint(idInt), username)
	if err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Invalid reservation id or not right permission to verify it"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.Response{Message: "Reservation successfully verified"})
}

func (uh *ReservationsHandler) GetAllByUserVerified(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	username, err := GetUsernameFromRequest(req)
	if err != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Unauthorized"})
		return
	}

	offset, size := uh.parseSearchPageable(req)
	reservations, totalElements, _ := uh.repository.FindAllByUser(username, true, offset, size)

	resDTOs := []models.ReservationDTO{}
	for _, reservation := range reservations {
		resDTOs = append(resDTOs, reservation.ToDTO())
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.Header().Set("total-elements", strconv.FormatInt(totalElements, 10))
	json.NewEncoder(resWriter).Encode(resDTOs)
}

func (uh *ReservationsHandler) GetAllByUserUnverified(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	username, err := GetUsernameFromRequest(req)
	if err != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Unauthorized"})
		return
	}

	offset, size := uh.parseSearchPageable(req)
	reservations, totalElements, _ := uh.repository.FindAllByUser(username, false, offset, size)

	resDTOs := []models.ReservationDTO{}
	for _, reservation := range reservations {
		resDTOs = append(resDTOs, reservation.ToDTO())
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.Header().Set("total-elements", strconv.FormatInt(totalElements, 10))
	json.NewEncoder(resWriter).Encode(resDTOs)
}

func (uh *ReservationsHandler) GetAllByDriverAndDriveVerified(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	driveIdStr := params["drive-id"]
	driveId, err := strconv.ParseInt(driveIdStr, 10, 64)

	username, err := GetUsernameFromRequest(req)
	if err != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Unauthorized"})
		return
	}

	offset, size := uh.parseSearchPageable(req)
	reservations, totalElements, _ := uh.repository.FindAllByDriverAndDriveId(username, driveId, true, offset, size)

	resDTOs := []models.ReservationDTO{}
	for _, reservation := range reservations {
		resDTOs = append(resDTOs, reservation.ToDTO())
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.Header().Set("total-elements", strconv.FormatInt(totalElements, 10))
	json.NewEncoder(resWriter).Encode(resDTOs)
}

func (uh *ReservationsHandler) GetAllByDriverAndDriveUnverified(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	driveIdStr := params["drive-id"]
	driveId, err := strconv.ParseInt(driveIdStr, 10, 64)

	username, err := GetUsernameFromRequest(req)
	if err != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Unauthorized"})
		return
	}

	offset, size := uh.parseSearchPageable(req)
	reservations, totalElements, _ := uh.repository.FindAllByDriverAndDriveId(username, driveId, false, offset, size)

	resDTOs := []models.ReservationDTO{}
	for _, reservation := range reservations {
		resDTOs = append(resDTOs, reservation.ToDTO())
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.Header().Set("total-elements", strconv.FormatInt(totalElements, 10))
	json.NewEncoder(resWriter).Encode(resDTOs)
}

func (uh *ReservationsHandler) GetAllByDriveIdVerified(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	driveIdStr := params["drive-id"]
	driveId, _ := strconv.ParseInt(driveIdStr, 10, 64)

	offset, size := uh.parseSearchPageable(req)
	reservations, totalElements, _ := uh.repository.FindAllByDriveId(driveId, true, offset, size)

	resDTOs := []models.ReservationDTO{}
	for _, reservation := range reservations {
		resDTOs = append(resDTOs, reservation.ToDTO())
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.Header().Set("total-elements", strconv.FormatInt(totalElements, 10))
	json.NewEncoder(resWriter).Encode(resDTOs)
}

func (uh *ReservationsHandler) DriveChanged(resWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	driveIdStr := params["drive-id"]
	driveId, _ := strconv.ParseInt(driveIdStr, 10, 64)

	reservations, _ := uh.repository.FindAllByDriveIdVerifUnverif(driveId)

	for _, reservation := range reservations {
		mailMessage := models.Mail{
			To:      reservation.GetPassengerEmail(),
			Subject: "Ride data changed",
			Body:    ComposeDriveChangedEmail(int(driveId), reservation.PassengerUsername),
		}
		uh.mailHandler.SendMail(mailMessage)
	}
}

func (uh *ReservationsHandler) IsVerifiedByDriveIdAndUser(resWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idStr := params["drive-id"]
	username := params["username"]
	idInt, _ := strconv.ParseInt(idStr, 10, 32)

	reservations, err := uh.repository.FindByDriveIdAndUsername(int32(idInt), username, true)
	if err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		return
	}

	for _, reservation := range reservations {
		if reservation.Verified {
			resWriter.WriteHeader(http.StatusOK)
			return
		}
	}

	resWriter.WriteHeader(http.StatusNotFound)
}

func (uh *ReservationsHandler) parseSearchPageable(req *http.Request) (int, int) {
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
