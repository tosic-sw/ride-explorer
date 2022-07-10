package db

import (
	"RatingService/models"
	"gorm.io/gorm"
)

var Ratings = []models.Rating{
	{
		Model:     gorm.Model{},
		Evaluator: "tica",
		Evaluated: "beli",
		DriveId:   1,
		Positive:  false,
		Text:      "Smoking in car",
	},
	{
		Evaluator: "tica",
		Evaluated: "beli",
		DriveId:   1,
		Positive:  false,
		Text:      "Sot smoking in car",
	},
}
