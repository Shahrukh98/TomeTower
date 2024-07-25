package service

import (
	"tometower/internal/dto"
	"tometower/internal/repository"
)

type AuthorService struct {
	repo repository.AuthorRepository
}

func NewAuthorService(repo repository.AuthorRepository) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) GetAllAuthors() ([]dto.AuthorDataDto, error) {
	authors, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var authorList []dto.AuthorDataDto
	for _, a := range authors {
		mappedAuthor, err := dto.MapAuthorToAuthorDataDto(a)
		if err != nil {
			return nil, err
		}
		authorList = append(authorList, mappedAuthor)
	}
	return authorList, nil
}

func (s *AuthorService) GetById(id string) (*dto.AuthorDataDto, error) {
	author, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	authorData, err := dto.MapAuthorToAuthorDataDto(author)
	if err != nil {
		return nil, err
	}
	return &authorData, nil
}

func (s *AuthorService) AddAuthor(createAuthorDto dto.CreateAuthorDto) (string, error) {
	author, err := dto.MapCreateAuthorDtoToAuthor(&createAuthorDto)
	if err != nil {
		return "", err
	}
	return s.repo.Add(*author)
}

func (s *AuthorService) RemoveAuthor(id string) error {
	return s.repo.Remove(id)
}
