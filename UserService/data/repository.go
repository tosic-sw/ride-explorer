package data

import (
	"UserService/models"
	"errors"
	"gorm.io/gorm"
	"strings"
	"time"
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

func (repo *Repository) FindOneAcc(username string) (*models.UserAccount, error) {
	var acc models.UserAccount

	result := repo.db.Where("username = ?", username).First(&acc)

	if result.Error != nil {
		return &acc, errors.New("account with given username does not exist")
	}

	return &acc, nil
}

func (repo *Repository) FindOneAccRole(username string) (*models.UserAccount, error) {
	var acc models.UserAccount

	result := repo.db.Where("username = ? AND verified = true", username).First(&acc)

	if result.Error != nil {
		return &acc, errors.New("account with given username does not exist")
	}

	return &acc, nil
}

func (repo *Repository) FindOneLogin(username string) (*models.UserAccount, error) {
	var acc models.UserAccount

	result := repo.db.Where("username = ? AND verified = true", username).First(&acc)

	if result.Error != nil {
		return &acc, errors.New("username does not exist in database")
	}

	if acc.BannedUntil > time.Now().UnixMilli() {
		return &acc, errors.New("banned until " + time.UnixMilli(acc.BannedUntil).Format("02.01.2006 15:05"))
	}

	return &acc, nil
}

func (repo *Repository) FindOneAdmin(username string) (*models.Admin, error) {
	var admin models.Admin

	result := repo.db.Where("username = ?", username).First(&admin)

	if result.Error != nil {
		return &admin, errors.New("username does not exists in database")
	}

	return &admin, nil
}

func (repo *Repository) FindOneDriver(username string) (*models.Driver, error) {
	var driver models.Driver

	result := repo.db.Where("username = ? AND verified = true", username).First(&driver)

	if result.Error != nil {
		return &driver, errors.New("username does not exists in database")
	}

	return &driver, nil
}

func (repo *Repository) FindOneDriverWithCar(username string) (*models.Driver, error) {
	var driver models.Driver

	result := repo.db.Preload("Car").Where("username = ? AND verified = true", username).First(&driver)

	if result.Error != nil {
		return &driver, errors.New("username does not exists in database")
	}

	return &driver, nil
}

func (repo *Repository) FindOneUnverifiedDriverWithCar(username string) (*models.Driver, error) {
	var driver models.Driver

	result := repo.db.Preload("Car").Where("username = ? AND verified = false", username).First(&driver)

	if result.Error != nil {
		return &driver, errors.New("not existent unverified driver for given username")
	}

	return &driver, nil
}

func (repo *Repository) FindOnePassenger(username string) (*models.Passenger, error) {
	var passenger models.Passenger

	result := repo.db.Where("username = ?", username).First(&passenger)

	if result.Error != nil {
		return &passenger, errors.New("username does not exists in database")
	}

	return &passenger, nil
}

func (repo *Repository) SaveUserAccount(acc *models.UserAccount) (*models.UserAccount, error) {
	result := repo.db.Create(acc)

	if result.Error != nil {
		return acc, result.Error
	}

	return acc, nil
}

func (repo *Repository) BanUserAccount(username string) (*models.UserAccount, error) {
	var acc models.UserAccount
	result := repo.db.Where("username = ?", username).First(&acc)

	if result.Error != nil {
		return &acc, result.Error
	}

	acc.BannedUntil = time.Now().AddDate(0, 3, 0).UnixMilli()
	repo.db.Save(acc)

	return &acc, nil
}

func (repo *Repository) DeleteUserAccount(username string) error {
	result := repo.db.Where("username = ?", username).Delete(&models.UserAccount{})

	return result.Error
}

func (repo *Repository) SaveAdmin(admin *models.Admin) (*models.Admin, error) {
	result := repo.db.Create(admin)

	if result.Error != nil {
		return admin, result.Error
	}

	return admin, nil
}

func (repo *Repository) SaveDriver(driver *models.Driver) (*models.Driver, error) {
	result := repo.db.Create(driver)

	if result.Error != nil {
		return driver, result.Error
	}

	return driver, nil
}

func (repo *Repository) VerifyDriver(username string) error {
	var driver models.Driver
	result := repo.db.Where("username = ?", username).First(&driver)
	if result.Error != nil {
		return result.Error
	}
	driver.Verified = true
	repo.db.Save(&driver)

	var account models.UserAccount
	result = repo.db.Where("username = ?", username).First(&account)
	if result.Error != nil {
		return result.Error
	}
	account.Verified = true
	repo.db.Save(&account)

	return nil
}

func (repo *Repository) SavePassenger(passenger *models.Passenger) (*models.Passenger, error) {
	result := repo.db.Create(passenger)

	if result.Error != nil {
		return passenger, result.Error
	}

	return passenger, nil
}

func (repo *Repository) UpdateAdmin(dto *models.UserForUpdateDTO, username string) (*models.Admin, error) {
	var admin models.Admin
	result := repo.db.Where("username = ?", username).First(&admin)

	if result.Error != nil {
		return &admin, result.Error
	}

	admin.Email = dto.Email
	admin.Firstname = dto.Firstname
	admin.Lastname = dto.Lastname
	repo.db.Save(&admin)

	return &admin, nil
}

func (repo *Repository) UpdateDriver(dto *models.UserForUpdateDTO, username string) (*models.Driver, error) {
	var driver models.Driver
	result := repo.db.Where("username = ? AND verified = true", username).First(&driver)

	if result.Error != nil {
		return &driver, errors.New("username does not exists in database or not verified")
	}

	driver.Email = dto.Email
	driver.Firstname = dto.Firstname
	driver.Lastname = dto.Lastname
	repo.db.Save(&driver)

	return &driver, nil
}

func (repo *Repository) UpdatePassenger(dto *models.UserForUpdateDTO, username string) (*models.Passenger, error) {
	var passenger models.Passenger
	result := repo.db.Where("username = ?", username).First(&passenger)

	if result.Error != nil {
		return &passenger, errors.New("username does not exists in database")
	}

	passenger.Email = dto.Email
	passenger.Firstname = dto.Firstname
	passenger.Lastname = dto.Lastname
	repo.db.Save(&passenger)

	return &passenger, nil
}

func (repo *Repository) ChangePassword(username string, password string) (*models.UserAccount, error) {
	var account models.UserAccount
	result := repo.db.Where("username = ?  AND verified = true", username).First(&account)

	if result.Error != nil {
		return &account, errors.New("username does not exists in database")
	}

	account.Password = password
	repo.db.Save(&account)

	return &account, nil
}

func (repo *Repository) DeleteDriver(username string) error {
	result := repo.db.Where("username = ?", username).Delete(&models.Driver{})

	return result.Error
}

func (repo *Repository) DeletePassenger(username string) error {
	result := repo.db.Where("username = ?", username).Delete(&models.Passenger{})

	return result.Error
}

func (repo *Repository) BanDriver(username string) (*models.Driver, *models.UserAccount, error) {
	var driver models.Driver
	var acc models.UserAccount

	result := repo.db.Where("username = ?", username).First(&driver)
	if result.Error != nil {
		return &driver, &acc, errors.New("username does not exists in database")
	}

	result = repo.db.Where("username = ?", username).First(&acc)
	if result.Error != nil {
		return &driver, &acc, errors.New("username does not exists in database")
	}

	driver.BannedUntil = time.Now().AddDate(0, 3, 0).UnixMilli()
	acc.BannedUntil = time.Now().AddDate(0, 3, 0).UnixMilli()

	repo.db.Save(&driver)
	repo.db.Save(&acc)

	return &driver, &acc, nil
}

func (repo *Repository) BanPassenger(username string) (*models.Passenger, *models.UserAccount, error) {
	var passenger models.Passenger
	var acc models.UserAccount

	result := repo.db.Where("username = ?", username).First(&passenger)
	if result.Error != nil {
		return &passenger, &acc, errors.New("username does not exists in database")
	}

	result = repo.db.Where("username = ?", username).First(&acc)
	if result.Error != nil {
		return &passenger, &acc, errors.New("username does not exists in database")
	}

	passenger.BannedUntil = time.Now().AddDate(0, 3, 0).UnixMilli()
	acc.BannedUntil = time.Now().AddDate(0, 3, 0).UnixMilli()

	repo.db.Save(&passenger)
	repo.db.Save(&acc)

	return &passenger, &acc, nil
}

func (repo *Repository) SearchAdmins(search string, offset int, size int) ([]*models.Admin, int64, error) {
	var admins []*models.Admin
	var totalElements int64 = -1

	result := repo.db.Scopes(repo.paginate(offset, size)).
		Where("(deleted_at IS NULL) AND "+
			"('' = ? or "+
			"firstname LIKE ? or "+
			"lastname LIKE ? or "+
			"email LIKE ? or "+
			"username LIKE ?)", search, concat(search), concat(search), concat(search), concat(search)).
		Find(&admins)

	result = repo.db.Table("admins").
		Where("(deleted_at IS NULL) AND "+
			"('' = ? or "+
			"firstname LIKE ? or "+
			"lastname LIKE ? or "+
			"email LIKE ? or "+
			"username LIKE ?)", search, concat(search), concat(search), concat(search), concat(search)).
		Count(&totalElements)

	if result.Error != nil {
		return admins, totalElements, result.Error
	}

	return admins, totalElements, nil
}
func (repo *Repository) SearchDrivers(search string, offset int, size int, verified bool) ([]*models.Driver, int64, error) {
	var drivers []*models.Driver
	var totalElements int64 = -1

	result := repo.db.Scopes(repo.paginate(offset, size)).
		Where("(deleted_at IS NULL AND banned_until < ? AND verified = ?) AND "+
			"('' = ? or "+
			"firstname LIKE ? or "+
			"lastname LIKE ? or "+
			"email LIKE ? or "+
			"username LIKE ?)", time.Now().UnixMilli(), verified, search, concat(search), concat(search), concat(search), concat(search)).
		Find(&drivers)

	result = repo.db.Table("drivers").
		Where("(deleted_at IS NULL AND banned_until < ? AND verified = ?) AND "+
			"('' = ? or "+
			"firstname LIKE ? or "+
			"lastname LIKE ? or "+
			"email LIKE ? or "+
			"username LIKE ?)", time.Now().UnixMilli(), verified, search, concat(search), concat(search), concat(search), concat(search)).
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
		Where("(deleted_at IS NULL AND banned_until < ?) AND "+
			"('' = ? or "+
			"firstname LIKE ? or "+
			"lastname LIKE ? or "+
			"email LIKE ? or "+
			"username LIKE ?)", time.Now().UnixMilli(), search, concat(search), concat(search), concat(search), concat(search)).
		Find(&passengers)

	result = repo.db.Table("passengers").
		Where("(deleted_at IS NULL AND banned_until < ?) AND "+
			"('' = ? or "+
			"firstname LIKE ? or "+
			"lastname LIKE ? or "+
			"email LIKE ? or "+
			"username LIKE ?)", time.Now().UnixMilli(), search, concat(search), concat(search), concat(search), concat(search)).
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
