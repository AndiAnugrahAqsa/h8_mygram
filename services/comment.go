package services

import (
	"mygram/dto"
	"mygram/repositories"
)

type CommentService struct {
	commentRepository repositories.CommentRepository
}

func NewCommentService(repository repositories.CommentRepository) CommentService {
	return CommentService{
		commentRepository: repository,
	}
}

func (s *CommentService) GetAll() (*[]dto.CommentResponse, error) {
	comments, err := s.commentRepository.GetAll()

	if err != nil {
		return nil, err
	}

	commentsResponse := []dto.CommentResponse{}

	for _, comment := range *comments {
		commentsResponse = append(commentsResponse, dto.FromCommentModelToResponse(comment))
	}

	return &commentsResponse, nil
}

func (s *CommentService) Create(commentRequest dto.CommentRequest) (*dto.CommentResponse, error) {
	comment, err := s.commentRepository.Create(commentRequest.ToModel())

	if err != nil {
		return nil, err
	}

	commentResponse := dto.FromCommentModelToCreateResponse(*comment)

	return &commentResponse, nil
}

func (s *CommentService) Update(id int, commentRequest dto.CommentRequest) (*dto.CommentResponse, error) {
	comment, err := s.commentRepository.Update(id, commentRequest.ToModel())

	if err != nil {
		return nil, err
	}

	commentResponse := dto.FromCommentModelToUpdateResponse(*comment)

	return &commentResponse, nil
}

func (s *CommentService) Delete(id int) error {
	err := s.commentRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
