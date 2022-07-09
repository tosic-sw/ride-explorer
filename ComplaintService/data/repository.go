package data

import (
	"ComplaintService/models"
	"errors"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) FindOne(id uint) (*models.Complaint, error) {
	var complaint models.Complaint

	result := repo.db.Where("id = ?", id).First(&complaint)

	if result.Error != nil {
		return nil, errors.New("complaint with given id does not exist")
	}

	return &complaint, nil
}

func (repo *Repository) FindOneComplex(accuser string, accused string, driveId uint) (*models.Complaint, error) {
	var complaint models.Complaint

	result := repo.db.Where("accuser = ? AND accused = ? AND drive_id = ?", accuser, accused, driveId).First(&complaint)

	if result.Error != nil {
		return nil, errors.New("complaint with given params does not exist")
	}

	return &complaint, nil
}

func (repo *Repository) SaveComplaint(complaint *models.Complaint) (*models.Complaint, error) {
	result := repo.db.Create(complaint)

	if result.Error != nil {
		return complaint, errors.New("an error occurred while saving complaint")
	}

	return complaint, nil
}

func (repo *Repository) DeleteComplaint(id uint) error {
	result := repo.db.Where("id = ?", id).Delete(&models.Complaint{})

	return result.Error
}

func (repo *Repository) FindAll(offset int, size int) ([]*models.Complaint, int64, error) {
	var complaints []*models.Complaint
	var totalElements int64 = -1

	result := repo.db.Scopes(repo.paginate(offset, size)).
		Order("id desc").
		Find(&complaints)

	result = repo.db.Table("complaints").
		Order("id desc").
		Find(&complaints).
		Count(&totalElements)

	if result.Error != nil {
		return complaints, totalElements, result.Error
	}

	return complaints, totalElements, nil
}

func (*Repository) paginate(offset int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(size)
	}
}
