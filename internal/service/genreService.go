package service

import (
	"tometower/internal/dto"
	"tometower/internal/entity"
	"tometower/internal/repository"
)

type GenreService struct {
	repo repository.GenreRepository
}

func NewGenreService(repo repository.GenreRepository) *GenreService {
	return &GenreService{repo: repo}
}

func (s *GenreService) GetAllGenres() ([]entity.Genre, error) {
	return s.repo.GetAll()
}

func (s *GenreService) GetGenreById(id string) (*dto.GenreDataDto, error) {
	genre, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	genreDataDto, err := dto.MapGenreToGenreDataDto(genre)
	if err != nil {
		return nil, err
	}
	return &genreDataDto, nil
}

func (s *GenreService) AddGenre(createGenreDto dto.CreateGenreDto) (string, error) {
	genre, err := dto.MapCreateGenreToGenre(createGenreDto)
	if err != nil {
		return "", err
	}
	return s.repo.Add(genre)
}

func (s *GenreService) RemoveGenre(id string) error {
	return s.repo.Remove(id)
}
