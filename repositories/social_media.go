package repositories

import (
	"mygram/models"

	"gorm.io/gorm"
)

func NewSocialMediaRepository(gormDB *gorm.DB) SocialMediaRepository {
	return &SocialMediaRepositoryImpl{
		db: gormDB,
	}
}

type SocialMediaRepositoryImpl struct {
	db *gorm.DB
}

func (r *SocialMediaRepositoryImpl) GetAll() (*[]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia

	err := r.db.Preload("User").Find(&socialMedias).Error
	if err != nil {
		return nil, err
	}

	return &socialMedias, nil
}

func (r *SocialMediaRepositoryImpl) GetOneByFilter(key string, value ...any) (*models.SocialMedia, error) {
	var socialMedia models.SocialMedia

	if err := r.db.Where(key, value...).First(&socialMedia).Error; err != nil {
		return nil, err
	}

	return &socialMedia, nil
}

func (r *SocialMediaRepositoryImpl) Create(socialMediaRequest models.SocialMedia) (*models.SocialMedia, error) {
	result := r.db.Create(&socialMediaRequest)

	if err := result.Error; err != nil {
		return nil, err
	}

	var socialMedia models.SocialMedia

	result.Last(&socialMedia)

	return &socialMedia, nil
}

func (r *SocialMediaRepositoryImpl) Update(id int, socialMediaRequest models.SocialMedia) (*models.SocialMedia, error) {
	socialMedia, err := r.GetOneByFilter("id", id)

	if err != nil {
		return nil, err
	}

	socialMedia.Name = socialMediaRequest.Name
	socialMedia.SocialMediaURL = socialMediaRequest.SocialMediaURL

	result := r.db.Save(socialMedia)

	if err := result.Error; err != nil {
		return nil, err
	}

	result.Last(&socialMedia)

	return socialMedia, nil
}

func (r *SocialMediaRepositoryImpl) Delete(id int) error {
	socialMedia, err := r.GetOneByFilter("id", id)

	if err != nil {
		return err
	}

	if err := r.db.Delete(&socialMedia).Error; err != nil {
		return err
	}

	return nil
}
