package dto

import (
	"mygram/models"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserRegister struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Age      int    `json:"age" validate:"required,min=8"`
}

func (u *UserRegister) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}

func (u *UserRegister) ToModel() models.User {
	return models.User{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Age:      u.Age,
	}
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (u *UserLogin) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}

func (u *UserLogin) ToModel() models.User {
	return models.User{
		Email:    u.Email,
		Password: u.Password,
	}
}

type UserEdit struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
}

func (u *UserEdit) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}

func (u *UserEdit) ToModel() models.User {
	return models.User{
		Username: u.Username,
		Email:    u.Email,
	}
}

type UserResponse struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Age       int        `json:"age"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func FromUserModelToResponse(user models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Age:       user.Age,
		CreatedAt: nil,
		UpdatedAt: nil,
	}
}

func FromUserModelToUpdateResponse(user models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Age:       user.Age,
		UpdatedAt: &user.UpdatedAt,
	}
}
