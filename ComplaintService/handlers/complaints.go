package handlers

import (
	"ComplaintService/data"
	"ComplaintService/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ComplaintsHandler struct {
	repository *data.Repository
}

func NewComplaintsHandler(repository *data.Repository) *ComplaintsHandler {
	return &ComplaintsHandler{repository}
}

func (ch *ComplaintsHandler) CreateComplaint(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)
	username, err := GetUsernameFromRequest(req)
	if err != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Unauthorized"})
		return
	}

	var complaintDTO models.CreateComplaintDTO
	err = json.NewDecoder(req.Body).Decode(&complaintDTO)
	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Invalid data sent"})
		return
	}

	_, err = ch.repository.FindOneComplex(username, complaintDTO.Accused, complaintDTO.DriveId)
	if err == nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Complaint already exists for drive"})
		return
	}

	role, err := GetRoleFromRequest(req)
	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: err.Error()})
		return
	}

	complaint := complaintDTO.ToComplaint(username)
	if role == "DRIVER" {
		ch.ProceedSavingAndReturnResponseDriver(&resWriter, &complaint)
	} else if role == "PASSENGER" {
		ch.ProceedSavingAndReturnResponsePassenger(&resWriter, &complaint)
	}
}

func (ch *ComplaintsHandler) ProceedSavingAndReturnResponseDriver(resWriter *http.ResponseWriter, complaint *models.Complaint) {
	if err := ExistsFinishedDriveDriver(complaint.DriveId, complaint.Accuser); err != nil {
		json.NewEncoder(*resWriter).Encode(models.Response{Message: err.Error()})
		return
	}

	if err := ExistsVerifiedReservation(complaint.DriveId, complaint.Accused); err != nil {
		json.NewEncoder(*resWriter).Encode(models.Response{Message: err.Error()})
		return
	}

	_, err := ch.repository.SaveComplaint(complaint)
	if err != nil {
		fmt.Println(err.Error())
		(*resWriter).WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(*resWriter).Encode(models.Response{Message: "Unknown error happened while saving complaint"})
		return
	}

	json.NewEncoder(*resWriter).Encode(models.Response{Message: "Complaint successfully created"})
}

func (ch *ComplaintsHandler) ProceedSavingAndReturnResponsePassenger(resWriter *http.ResponseWriter, complaint *models.Complaint) {
	accusedRole, err := GetRoleOfUser(complaint.Accused)
	if err != nil {
		(*resWriter).WriteHeader(http.StatusBadRequest)
		json.NewEncoder(*resWriter).Encode(models.Response{Message: err.Error()})
		return
	}

	if accusedRole == "DRIVER" {
		if err := ExistsFinishedDriveDriver(complaint.DriveId, complaint.Accused); err != nil {
			(*resWriter).WriteHeader(http.StatusBadRequest)
			json.NewEncoder(*resWriter).Encode(models.Response{Message: err.Error()})
			return
		}
		if err := ExistsVerifiedReservation(complaint.DriveId, complaint.Accuser); err != nil {
			(*resWriter).WriteHeader(http.StatusBadRequest)
			json.NewEncoder(*resWriter).Encode(models.Response{Message: err.Error()})
			return
		}
	} else if accusedRole == "PASSENGER" {
		if err := ExistsFinishedDrive(complaint.DriveId); err != nil {
			(*resWriter).WriteHeader(http.StatusBadRequest)
			json.NewEncoder(*resWriter).Encode(models.Response{Message: err.Error()})
			return
		}
		if err := ExistsVerifiedReservation(complaint.DriveId, complaint.Accuser); err != nil {
			(*resWriter).WriteHeader(http.StatusBadRequest)
			json.NewEncoder(*resWriter).Encode(models.Response{Message: err.Error()})
			return
		}
		if err := ExistsVerifiedReservation(complaint.DriveId, complaint.Accused); err != nil {
			(*resWriter).WriteHeader(http.StatusBadRequest)
			json.NewEncoder(*resWriter).Encode(models.Response{Message: err.Error()})
			return
		}
	}

	_, err = ch.repository.SaveComplaint(complaint)
	if err != nil {
		fmt.Println(err.Error())
		(*resWriter).WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(*resWriter).Encode(models.Response{Message: "Unknown error happened while saving complaint"})
		return
	}

	json.NewEncoder(*resWriter).Encode(models.Response{Message: "Complaint successfully created"})
}

func (ch *ComplaintsHandler) DeleteComplaint(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, err := strconv.ParseInt(idStr, 10, 64)

	if _, err := ch.repository.FindOne(uint(idInt)); err == nil {
		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Invalid complaint id sent for deletion"})
		return
	}

	err = ch.repository.DeleteComplaint(uint(idInt))
	if err != nil {
		fmt.Println(err.Error())
		resWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Unknown error happened while deleting complaint"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.Response{Message: "Complaint successfully deleted"})
}

func (ch *ComplaintsHandler) GetAllComplaints(resWriter http.ResponseWriter, req *http.Request) {
	offset, size := ch.parseSearchPageable(req)

	complaints, totalElements, _ := ch.repository.FindAll(offset, size)
	resDTOs := []models.ComplaintDTO{}

	for _, complaint := range complaints {
		resDTOs = append(resDTOs, complaint.ToDTO())
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.Header().Set("total-elements", strconv.FormatInt(totalElements, 10))
	json.NewEncoder(resWriter).Encode(resDTOs)
}

func (ch *ComplaintsHandler) parseSearchPageable(req *http.Request) (int, int) {
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
