package repositories

import (
	"mygram/models"

	"gorm.io/gorm"
)

func NewPhotoRepository(gormDB *gorm.DB) PhotoRepository {
	return &PhotoRepositoryImpl{
		db: gormDB,
	}
}

type PhotoRepositoryImpl struct {
	db *gorm.DB
}

func (r *PhotoRepositoryImpl) GetAll() (*[]models.Photo, error) {
	var photos []models.Photo

	err := r.db.Preload("User").Find(&photos).Error
	if err != nil {
		return nil, err
	}

	return &photos, nil
}

func (r *PhotoRepositoryImpl) GetOneByFilter(key string, value ...any) (*models.Photo, error) {
	var photo models.Photo

	if err := r.db.Where(key, value...).First(&photo).Error; err != nil {
		return nil, err
	}

	return &photo, nil
}

func (r *PhotoRepositoryImpl) Create(photoRequest models.Photo) (*models.Photo, error) {
	result := r.db.Create(&photoRequest)

	if err := result.Error; err != nil {
		return nil, err
	}

	var photo models.Photo

	result.Last(&photo)

	return &photo, nil
}

func (r *PhotoRepositoryImpl) Update(id int, photoRequest models.Photo) (*models.Photo, error) {
	photo, err := r.GetOneByFilter("id", id)

	if err != nil {
		return nil, err
	}

	photo.Title = photoRequest.Title
	photo.Caption = photoRequest.Caption
	photo.PhotoURL = photoRequest.PhotoURL

	result := r.db.Save(photo)

	if err := result.Error; err != nil {
		return nil, err
	}

	result.Last(&photo)

	return photo, nil
}

func (r *PhotoRepositoryImpl) Delete(id int) error {
	photo, err := r.GetOneByFilter("id", id)

	if err != nil {
		return err
	}

	if err := r.db.Select("Comments").Delete(&photo).Error; err != nil {
		return err
	}

	return nil
}
