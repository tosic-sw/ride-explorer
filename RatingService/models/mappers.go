package models

import "gorm.io/gorm"

func (rating *Rating) ToDTO() ViewRatingDTO {
	return ViewRatingDTO{
		Id:        rating.ID,
		Evaluator: rating.Evaluator,
		Evaluated: rating.Evaluated,
		DriveId:   rating.DriveId,
		Positive:  rating.Positive,
		Text:      rating.Text,
	}
}

func (dto *RatingDTO) ToRating(evaluator string) Rating {
	return Rating{
		Model:     gorm.Model{},
		Evaluator: evaluator,
		Evaluated: dto.Evaluated,
		DriveId:   dto.DriveId,
		Positive:  dto.Positive,
		Text:      dto.Text,
	}
}
