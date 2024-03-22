package services

import (
	"mygram/dto"
	"mygram/repositories"
)

type SocialMediaService struct {
	socialMediaRepository repositories.SocialMediaRepository
}

func NewSocialMediaService(repository repositories.SocialMediaRepository) SocialMediaService {
	return SocialMediaService{
		socialMediaRepository: repository,
	}
}

func (s *SocialMediaService) GetAll() (*[]dto.SocialMediaResponse, error) {
	socialMedias, err := s.socialMediaRepository.GetAll()

	if err != nil {
		return nil, err
	}

	socialMediasResponse := []dto.SocialMediaResponse{}

	for _, socialMedia := range *socialMedias {
		socialMediasResponse = append(socialMediasResponse, dto.FromSocialMediaModelToResponse(socialMedia))
	}

	return &socialMediasResponse, nil
}

func (s *SocialMediaService) Create(socialMediaRequest dto.SocialMediaRequest) (*dto.SocialMediaResponse, error) {
	socialMedia, err := s.socialMediaRepository.Create(socialMediaRequest.ToModel())

	if err != nil {
		return nil, err
	}

	socialMediaResponse := dto.FromSocialMediaModelToCreateResponse(*socialMedia)

	return &socialMediaResponse, nil
}

func (s *SocialMediaService) Update(id int, socialMediaRequest dto.SocialMediaRequest) (*dto.SocialMediaResponse, error) {
	socialMedia, err := s.socialMediaRepository.Update(id, socialMediaRequest.ToModel())

	if err != nil {
		return nil, err
	}

	socialMediaResponse := dto.FromSocialMediaModelToUpdateResponse(*socialMedia)

	return &socialMediaResponse, nil
}

func (s *SocialMediaService) Delete(id int) error {
	err := s.socialMediaRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
