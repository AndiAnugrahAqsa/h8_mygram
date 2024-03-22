package dto

import (
	"mygram/models"
	"time"

	"github.com/go-playground/validator/v10"
)

type CommentRequest struct {
	Message string `json:"message" validate:"required"`
	PhotoID int    `json:"photo_id"`
	UserID  int
}

func (u *CommentRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}

func (u *CommentRequest) ToModel() models.Comment {
	return models.Comment{
		Message: u.Message,
		PhotoID: u.PhotoID,
		UserID:  u.UserID,
	}
}

type CommentResponse struct {
	ID        int            `json:"id"`
	Message   string         `json:"message"`
	PhotoID   int            `json:"photo_id"`
	UserID    int            `json:"user_id"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	User      *UserResponse  `json:"user,omitempty"`
	Photo     *PhotoResponse `json:"photo,omitempty"`
}

func FromCommentModelToResponse(comment models.Comment) CommentResponse {
	return CommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: &comment.CreatedAt,
		UpdatedAt: &comment.UpdatedAt,
		User: &UserResponse{
			ID:       comment.User.ID,
			Email:    comment.User.Email,
			Username: comment.User.Username,
		},
		Photo: &PhotoResponse{
			ID:       comment.Photo.ID,
			Title:    comment.Photo.Title,
			Caption:  comment.Photo.Caption,
			PhotoURL: comment.Photo.PhotoURL,
			UserID:   comment.Photo.UserID,
		},
	}
}

func FromCommentModelToUpdateResponse(comment models.Comment) CommentResponse {
	return CommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: nil,
		UpdatedAt: &comment.UpdatedAt,
	}
}

func FromCommentModelToCreateResponse(comment models.Comment) CommentResponse {
	return CommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: &comment.CreatedAt,
		UpdatedAt: nil,
	}
}
