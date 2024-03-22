package services

import (
	"mygram/dto"
	"mygram/repositories"
)

type PhotoService struct {
	photoRepository repositories.PhotoRepository
}

func NewPhotoService(repository repositories.PhotoRepository) PhotoService {
	return PhotoService{
		photoRepository: repository,
	}
}

func (s *PhotoService) GetAll() (*[]dto.PhotoResponse, error) {
	photos, err := s.photoRepository.GetAll()

	if err != nil {
		return nil, err
	}

	photosResponse := []dto.PhotoResponse{}

	for _, photo := range *photos {
		photosResponse = append(photosResponse, dto.FromPhotoModelToResponse(photo))
	}

	return &photosResponse, nil
}

func (s *PhotoService) Create(photoRequest dto.PhotoRequest) (*dto.PhotoResponse, error) {
	photo, err := s.photoRepository.Create(photoRequest.ToModel())

	if err != nil {
		return nil, err
	}

	photoResponse := dto.FromPhotoModelToCreateResponse(*photo)

	return &photoResponse, nil
}

func (s *PhotoService) Update(id int, photoRequest dto.PhotoRequest) (*dto.PhotoResponse, error) {
	photo, err := s.photoRepository.Update(id, photoRequest.ToModel())

	if err != nil {
		return nil, err
	}

	photoResponse := dto.FromPhotoModelToUpdateResponse(*photo)

	return &photoResponse, nil
}

func (s *PhotoService) Delete(id int) error {
	err := s.photoRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
