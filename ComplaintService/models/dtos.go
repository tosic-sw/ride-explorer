package models

type ComplaintDTO struct {
	Accuser   string `json:"accuser"`
	Accused   string `json:"accused"`
	DriveId   uint   `json:"driveId"`
	Text      string `json:"text"`
	CreatedAt int64  `json:"createdAt"`
	Id        uint   `json:"id"`
}

type CreateComplaintDTO struct {
	Accused string `json:"accused"`
	DriveId uint   `json:"driveId"`
	Text    string `json:"text"`
}

type Response struct {
	Message string `json:"message"`
}

type RoleDTO struct {
	Role string `json:"role"`
}
