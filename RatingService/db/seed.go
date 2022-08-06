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
		Evaluator: "beli",
		Evaluated: "tica",
		DriveId:   1,
		Positive:  false,
		Text:      "Not smoking in car",
	},
}
