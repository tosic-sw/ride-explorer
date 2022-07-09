package models

import "gorm.io/gorm"

func (complaint *Complaint) ToDTO() ComplaintDTO {
	return ComplaintDTO{
		Accuser:   complaint.Accuser,
		Accused:   complaint.Accused,
		DriveId:   complaint.DriveId,
		Text:      complaint.Text,
		CreatedAt: complaint.CreatedAt.UnixMilli(),
		Id:        complaint.ID,
	}
}

func (dto *CreateComplaintDTO) ToComplaint(accuser string) Complaint {
	return Complaint{
		Model:   gorm.Model{},
		Accuser: accuser,
		Accused: dto.Accused,
		DriveId: dto.DriveId,
		Text:    dto.Text,
	}
}
