package db

import (
	"ComplaintService/models"
	"gorm.io/gorm"
)

var Complaints = []models.Complaint{
	{
		Model:   gorm.Model{},
		Accuser: "tica",
		Accused: "beli",
		DriveId: 0,
		Text:    "Smoking in car",
	},
	{
		Model:   gorm.Model{},
		Accuser: "beli",
		Accused: "tica",
		DriveId: 0,
		Text:    "Not smoking in car",
	},
}
