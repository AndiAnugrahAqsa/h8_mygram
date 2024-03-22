package dto

import (
	"mygram/models"
	"time"

	"github.com/go-playground/validator/v10"
)

type PhotoRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" validate:"required"`
	UserID   int
}

func (u *PhotoRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}

func (u *PhotoRequest) ToModel() models.Photo {
	return models.Photo{
		Title:    u.Title,
		Caption:  u.Caption,
		PhotoURL: u.PhotoURL,
		UserID:   u.UserID,
	}
}

type PhotoResponse struct {
	ID        int           `json:"id"`
	Title     string        `json:"title"`
	Caption   string        `json:"caption"`
	PhotoURL  string        `json:"photo_url"`
	UserID    int           `json:"user_id"`
	CreatedAt *time.Time    `json:"created_at,omitempty"`
	UpdatedAt *time.Time    `json:"updated_at,omitempty"`
	User      *UserResponse `json:"user,omitempty"`
}

func FromPhotoModelToResponse(photo models.Photo) PhotoResponse {
	return PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: &photo.CreatedAt,
		UpdatedAt: &photo.UpdatedAt,
		User: &UserResponse{
			Email:    photo.User.Email,
			Username: photo.User.Username,
		},
	}
}

func FromPhotoModelToUpdateResponse(photo models.Photo) PhotoResponse {
	return PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: nil,
		UpdatedAt: &photo.UpdatedAt,
		User:      nil,
	}
}

func FromPhotoModelToCreateResponse(photo models.Photo) PhotoResponse {
	return PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: &photo.CreatedAt,
		UpdatedAt: nil,
		User:      nil,
	}
}
