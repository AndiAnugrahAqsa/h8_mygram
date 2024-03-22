package repositories

import "mygram/models"

type UserRepository interface {
	GetOneByFilter(key string, value ...any) (*models.User, error)
	Create(userRequest models.User) (*models.User, error)
	Update(id int, userRequest models.User) (*models.User, error)
	Delete(id int) error
}

type PhotoRepository interface {
	GetAll() (*[]models.Photo, error)
	GetOneByFilter(key string, value ...any) (*models.Photo, error)
	Create(userRequest models.Photo) (*models.Photo, error)
	Update(id int, userRequest models.Photo) (*models.Photo, error)
	Delete(id int) error
}

type CommentRepository interface {
	GetAll() (*[]models.Comment, error)
	GetOneByFilter(key string, value ...any) (*models.Comment, error)
	Create(userRequest models.Comment) (*models.Comment, error)
	Update(id int, userRequest models.Comment) (*models.Comment, error)
	Delete(id int) error
}

type SocialMediaRepository interface {
	GetAll() (*[]models.SocialMedia, error)
	GetOneByFilter(key string, value ...any) (*models.SocialMedia, error)
	Create(userRequest models.SocialMedia) (*models.SocialMedia, error)
	Update(id int, userRequest models.SocialMedia) (*models.SocialMedia, error)
	Delete(id int) error
}
