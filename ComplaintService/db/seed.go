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
		DriveId: 1,
		Text:    "Smoking in car",
	},
	{
		Model:   gorm.Model{},
		Accuser: "beli",
		Accused: "tica",
		DriveId: 1,
		Text:    "Not smoking in car",
	},
}
