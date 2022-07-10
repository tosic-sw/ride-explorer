package data

import (
	"RatingService/models"
	"errors"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) FindOne(id uint) (*models.Rating, error) {
	var rating models.Rating

	result := repo.db.Where("id = ?", id).First(&rating)

	if result.Error != nil {
		return nil, errors.New("rating with given id does not exist")
	}

	return &rating, nil
}

func (repo *Repository) FindOneByEvaluator(id uint, evaluator string) (*models.Rating, error) {
	var rating models.Rating

	result := repo.db.Where("id = ? AND evaluator = ?", id, evaluator).First(&rating)

	if result.Error != nil {
		return nil, errors.New("invalid rating id and evaluator combination")
	}

	return &rating, nil
}

func (repo *Repository) FindOneComplex(evaluator string, evaluated string, driveId uint) (*models.Rating, error) {
	var rating models.Rating

	result := repo.db.Where("evaluator = ? AND evaluated = ? AND drive_id = ?", evaluator, evaluated, driveId).First(&rating)

	if result.Error != nil {
		return nil, errors.New("rating with given params does not exist")
	}

	return &rating, nil
}

func (repo *Repository) SaveRating(rating *models.Rating) (*models.Rating, error) {
	result := repo.db.Create(rating)

	if result.Error != nil {
		return rating, errors.New("an error occurred while saving rating")
	}

	return rating, nil
}

func (repo *Repository) UpdateRating(id uint, evaluator string, positive bool, text string) (*models.Rating, error) {
	var rating *models.Rating

	result := repo.db.Where("id = ? AND evaluator = ?", id, evaluator).First(&rating)

	if result.Error != nil {
		return nil, errors.New("invalid rating id and evaluator combination")
	}

	rating.Positive = positive
	rating.Text = text
	result = repo.db.Save(rating)

	if result.Error != nil {
		return nil, result.Error
	}

	return rating, nil
}

func (repo *Repository) DeleteRating(id uint) error {
	result := repo.db.Where("id = ?", id).Delete(&models.Rating{})

	return result.Error
}

func (repo *Repository) FindAllForEvaluated(evaluated string, offset int, size int) ([]*models.Rating, int64, error) {
	var ratings []*models.Rating
	var totalElements int64 = -1

	result := repo.db.Scopes(repo.paginate(offset, size)).
		Where("evaluated = ?", evaluated).
		Order("id desc").
		Find(&ratings)

	result = repo.db.Table("ratings").
		Where("evaluated = ?", evaluated).
		Order("id desc").
		Count(&totalElements)

	if result.Error != nil {
		return ratings, totalElements, result.Error
	}

	return ratings, totalElements, nil
}

func (*Repository) paginate(offset int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(size)
	}
}
