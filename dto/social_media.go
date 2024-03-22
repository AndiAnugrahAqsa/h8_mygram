package dto

import (
	"mygram/models"
	"time"

	"github.com/go-playground/validator/v10"
)

type SocialMediaRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaURL string `json:"social_media_url" validate:"required"`
	UserID         int
}

func (u *SocialMediaRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}

func (u *SocialMediaRequest) ToModel() models.SocialMedia {
	return models.SocialMedia{
		Name:           u.Name,
		SocialMediaURL: u.SocialMediaURL,
		UserID:         u.UserID,
	}
}

type SocialMediaResponse struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	SocialMediaURL string        `json:"social_media_url"`
	UserID         int           `json:"user_id"`
	CreatedAt      *time.Time    `json:"created_at,omitempty"`
	UpdatedAt      *time.Time    `json:"updated_at,omitempty"`
	User           *UserResponse `json:"user,omitempty"`
}

func FromSocialMediaModelToResponse(socialMedia models.SocialMedia) SocialMediaResponse {
	return SocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
		CreatedAt:      &socialMedia.CreatedAt,
		UpdatedAt:      &socialMedia.UpdatedAt,
		User: &UserResponse{
			ID:       socialMedia.User.ID,
			Email:    socialMedia.User.Email,
			Username: socialMedia.User.Username,
		},
	}
}

func FromSocialMediaModelToUpdateResponse(socialMedia models.SocialMedia) SocialMediaResponse {
	return SocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
		CreatedAt:      nil,
		UpdatedAt:      &socialMedia.UpdatedAt,
	}
}

func FromSocialMediaModelToCreateResponse(socialMedia models.SocialMedia) SocialMediaResponse {
	return SocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
		CreatedAt:      &socialMedia.CreatedAt,
		UpdatedAt:      nil,
	}
}
