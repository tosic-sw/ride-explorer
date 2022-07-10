package models

type RoleDTO struct {
	Role string `json:"role"`
}

type RatingDTO struct {
	Evaluated string `json:"evaluated"`
	DriveId   uint   `json:"driveId"`
	Positive  bool   `json:"positive"`
	Text      string `json:"text"`
}

type ViewRatingDTO struct {
	Id        uint   `json:"id"`
	Evaluator string `json:"evaluator"`
	Evaluated string `json:"evaluated"`
	DriveId   uint   `json:"driveId"`
	Positive  bool   `json:"positive"`
	Text      string `json:"text"`
}

type Response struct {
	Message string `json:"message"`
}
