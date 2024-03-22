package services

import (
	"mygram/dto"
	"mygram/models"
	"mygram/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return UserService{
		userRepository: repository,
	}
}

func (s *UserService) Create(userRequest dto.UserRegister) (*dto.UserResponse, error) {
	user, err := s.userRepository.Create(userRequest.ToModel())

	if err != nil {
		return nil, err
	}

	userResponse := dto.FromUserModelToResponse(*user)

	return &userResponse, nil
}

func (s *UserService) GetOneByFilter(key string, value ...any) (*models.User, error) {
	user, err := s.userRepository.GetOneByFilter(key, value...)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Update(id int, userRequest dto.UserEdit) (*dto.UserResponse, error) {
	user, err := s.userRepository.Update(id, userRequest.ToModel())

	if err != nil {
		return nil, err
	}

	userResponse := dto.FromUserModelToUpdateResponse(*user)

	return &userResponse, nil
}

func (s *UserService) Delete(id int) error {
	err := s.userRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
