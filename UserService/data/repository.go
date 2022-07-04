package data

import (
	"UserService/models"
	"gorm.io/gorm"
	"strings"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func concat(str string) string {
	return "%" + strings.ToLower(str) + "%"
}

func (repo *Repository) FindOneAdmin(username string) (*models.Admin, error) {
	var admin models.Admin

	result := repo.db.Where("username = ?", username).First(&admin)

	if result.Error != nil {
		return &admin, result.Error
	}

	return &admin, nil
}

func (repo *Repository) FindOneDriver(username string) (*models.Driver, error) {
	var driver models.Driver

	result := repo.db.Where("username = ?", username).First(&driver)

	if result.Error != nil {
		return &driver, result.Error
	}

	return &driver, nil
}

func (repo *Repository) FindOnePassenger(username string) (*models.Passenger, error) {
	var passenger models.Passenger

	result := repo.db.Where("username = ?", username).First(&passenger)

	if result.Error != nil {
		return &passenger, result.Error
	}

	return &passenger, nil
}

func (repo *Repository) saveAdmin(admin *models.Admin) (*models.Admin, error) {
	result := repo.db.Create(admin)

	if result.Error != nil {
		return admin, result.Error
	}

	return admin, nil
}

func (repo *Repository) saveDriver(driver *models.Driver) (*models.Driver, error) {
	result := repo.db.Create(driver)

	if result.Error != nil {
		return driver, result.Error
	}

	return driver, nil
}

func (repo *Repository) savePassenger(passenger *models.Passenger) (*models.Passenger, error) {
	result := repo.db.Create(passenger)

	if result.Error != nil {
		return passenger, result.Error
	}

	return passenger, nil
}

func (repo *Repository) SearchAdmins(search string, offset int, size int) ([]*models.Admin, int64, error) {
	var admins []*models.Admin
	var totalElements int64 = -1

	result := repo.db.Scopes(repo.paginate(offset, size)).
		Where("deleted_at IS NULL AND '' = ?", search).
		Or("email LIKE ?", concat(search)).
		Or("username LIKE ?", concat(search)).
		Or("firstname LIKE ?", concat(search)).
		Or("lastname LIKE ?", concat(search)).
		Find(&admins)

	result = repo.db.Table("admins").
		Where("deleted_at IS NULL AND '' = ?", search).
		Or("email LIKE ?", concat(search)).
		Or("username LIKE ?", concat(search)).
		Or("firstname LIKE ?", concat(search)).
		Or("lastname LIKE ?", concat(search)).
		Count(&totalElements)

	if result.Error != nil {
		return admins, totalElements, result.Error
	}

	return admins, totalElements, nil
}
func (repo *Repository) SearchDrivers(search string, offset int, size int) ([]*models.Driver, int64, error) {
	var drivers []*models.Driver
	var totalElements int64 = -1

	result := repo.db.Scopes(repo.paginate(offset, size)).
		Where("deleted_at IS NULL AND '' = ?", search).
		Or("email LIKE ?", concat(search)).
		Or("username LIKE ?", concat(search)).
		Or("firstname LIKE ?", concat(search)).
		Or("lastname LIKE ?", concat(search)).
		Find(&drivers)

	result = repo.db.Table("drivers").
		Where("deleted_at IS NULL AND '' = ?", search).
		Or("email LIKE ?", concat(search)).
		Or("username LIKE ?", concat(search)).
		Or("firstname LIKE ?", concat(search)).
		Or("lastname LIKE ?", concat(search)).
		Count(&totalElements)

	if result.Error != nil {
		return drivers, totalElements, result.Error
	}

	return drivers, totalElements, nil
}

func (repo *Repository) SearchPassengers(search string, offset int, size int) ([]*models.Passenger, int64, error) {
	var passengers []*models.Passenger
	var totalElements int64 = -1

	result := repo.db.Scopes(repo.paginate(offset, size)).
		Where("deleted_at IS NULL AND '' = ?", search).
		Or("email LIKE ?", concat(search)).
		Or("username LIKE ?", concat(search)).
		Or("firstname LIKE ?", concat(search)).
		Or("lastname LIKE ?", concat(search)).
		Find(&passengers)

	result = repo.db.Table("passengers").
		Where("deleted_at IS NULL AND '' = ?", search).
		Or("email LIKE ?", concat(search)).
		Or("username LIKE ?", concat(search)).
		Or("firstname LIKE ?", concat(search)).
		Or("lastname LIKE ?", concat(search)).
		Count(&totalElements)

	if result.Error != nil {
		return passengers, totalElements, result.Error
	}

	return passengers, totalElements, nil
}

func (*Repository) paginate(offset int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(size)
	}
}
