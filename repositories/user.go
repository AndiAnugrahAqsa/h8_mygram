package repositories

import (
	"mygram/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewUserRepository(gormDB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: gormDB,
	}
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (r *UserRepositoryImpl) GetOneByFilter(key string, value ...any) (*models.User, error) {
	var user models.User

	if err := r.db.Where(key, value...).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) Create(userRequest models.User) (*models.User, error) {
	result := r.db.Create(&userRequest)

	if err := result.Error; err != nil {
		return nil, err
	}

	var user models.User

	result.Last(&user)

	return &user, nil
}

func (r *UserRepositoryImpl) Update(id int, userRequest models.User) (*models.User, error) {
	user, err := r.GetOneByFilter("id", id)

	if err != nil {
		return nil, err
	}

	user.Email = userRequest.Email
	user.Username = userRequest.Username

	result := r.db.Save(user)

	if err := result.Error; err != nil {
		return nil, err
	}

	result.Last(user)

	return user, nil
}

func (r *UserRepositoryImpl) Delete(id int) error {
	user, err := r.GetOneByFilter("id", id)

	if err != nil {
		return err
	}
	photos := []models.Photo{}

	if err := r.db.Find(&photos, "user_id", user.ID).Error; err != nil {
		return err
	}

	if len(photos) != 0 {
		if err := r.db.Select(clause.Associations).Delete(&photos).Error; err != nil {
			return err
		}
	}

	if err := r.db.Select(clause.Associations).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
