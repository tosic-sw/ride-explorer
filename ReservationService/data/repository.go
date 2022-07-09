package data

import (
	"ReservationService/models"
	"errors"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) FindOne(id uint) (*models.Reservation, error) {
	var res models.Reservation

	result := repo.db.Where("id = ?", id).First(&res)

	if result.Error != nil {
		return nil, errors.New("reservation with given id does not exist")
	}

	return &res, nil
}

func (repo *Repository) FindOneByUser(id uint, username string) (*models.Reservation, error) {
	var res models.Reservation

	result := repo.db.Where("id = ? AND passenger_username = ?", id, username).First(&res)

	if result.Error != nil {
		return nil, errors.New("reservation with given id does not exist or you do not have permission to access it")
	}

	return &res, nil
}

func (repo *Repository) FindByDriveIdAndUsername(driveId int32, username string, verified bool) ([]*models.Reservation, error) {
	var reservations []*models.Reservation

	result := repo.db.Where("drive_id = ? AND passenger_username = ? AND verified = ?", driveId, username, verified).
		Find(&reservations)

	if result.Error != nil {
		return nil, result.Error
	}

	return reservations, nil
}

func (repo *Repository) SaveReservation(res *models.Reservation) (*models.Reservation, error) {
	result := repo.db.Create(res)

	if result.Error != nil {
		return res, errors.New("an error occurred while saving reservation")
	}

	return res, nil
}

func (repo *Repository) DeleteReservation(id uint, username string) error {
	result := repo.db.Where("id = ? AND passenger_username = ?", id, username).Delete(&models.Reservation{})

	return result.Error
}

func (repo *Repository) VerifyReservation(id uint, username string) (*models.Reservation, error) {
	var res models.Reservation
	result := repo.db.Where("id = ? AND driver_username = ?", id, username).First(&res)

	if result.Error != nil {
		return nil, result.Error
	}

	res.Verified = true
	repo.db.Save(&res)

	return &res, nil
}

func (repo *Repository) FindAllByUser(username string, verified bool, offset int, size int) ([]*models.Reservation, int64, error) {
	var reservations []*models.Reservation
	var totalElements int64 = -1

	result := repo.db.Scopes(repo.paginate(offset, size)).
		Where("(deleted_at IS NULL) AND passenger_username = ? AND verified = ?", username, verified).
		Find(&reservations)

	result = repo.db.Table("reservations").
		Where("(deleted_at IS NULL) AND passenger_username = ? AND verified = ?", username, verified).
		Count(&totalElements)

	if result.Error != nil {
		return reservations, totalElements, result.Error
	}

	return reservations, totalElements, nil
}

func (*Repository) paginate(offset int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(size)
	}
}
