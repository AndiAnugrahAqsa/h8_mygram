package repositories

import (
	"mygram/models"

	"gorm.io/gorm"
)

func NewCommentRepository(gormDB *gorm.DB) CommentRepository {
	return &CommentRepositoryImpl{
		db: gormDB,
	}
}

type CommentRepositoryImpl struct {
	db *gorm.DB
}

func (r *CommentRepositoryImpl) GetAll() (*[]models.Comment, error) {
	var comments []models.Comment

	err := r.db.Preload("User").Preload("Photo").Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return &comments, nil
}

func (r *CommentRepositoryImpl) GetOneByFilter(key string, value ...any) (*models.Comment, error) {
	var comment models.Comment

	if err := r.db.Where(key, value...).First(&comment).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *CommentRepositoryImpl) Create(commentRequest models.Comment) (*models.Comment, error) {
	result := r.db.Create(&commentRequest)

	if err := result.Error; err != nil {
		return nil, err
	}

	var comment models.Comment

	result.Last(&comment)

	return &comment, nil
}

func (r *CommentRepositoryImpl) Update(id int, commentRequest models.Comment) (*models.Comment, error) {
	comment, err := r.GetOneByFilter("id", id)

	if err != nil {
		return nil, err
	}

	comment.Message = commentRequest.Message

	result := r.db.Save(comment)

	if err := result.Error; err != nil {
		return nil, err
	}

	result.Last(&comment)

	return comment, nil
}

func (r *CommentRepositoryImpl) Delete(id int) error {
	comment, err := r.GetOneByFilter("id", id)

	if err != nil {
		return err
	}

	if err := r.db.Delete(&comment).Error; err != nil {
		return err
	}

	return nil
}
