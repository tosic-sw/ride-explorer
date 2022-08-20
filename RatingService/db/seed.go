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
		DriveId:   500,
		Positive:  false,
		Text:      "Smoking in car",
	},
	{
		Evaluator: "beli",
		Evaluated: "tica",
		DriveId:   500,
		Positive:  false,
		Text:      "Not smoking in car.",
	},
	{
		Evaluator: "mile",
		Evaluated: "ukica",
		DriveId:   501,
		Positive:  true,
		Text:      "Drives very slow but safe.",
	},
	{
		Evaluator: "ukica",
		Evaluated: "mile",
		DriveId:   501,
		Positive:  true,
		Text:      "All time complaining because of slow ride but good guy.",
	},
	{
		Evaluator: "toma",
		Evaluated: "mile",
		DriveId:   501,
		Positive:  false,
		Text:      "Want to be in middle of attention always. Loud talking.",
	},
	{
		Evaluator: "mile",
		Evaluated: "toma",
		DriveId:   501,
		Positive:  false,
		Text:      "Not a funny guy.",
	},
}
