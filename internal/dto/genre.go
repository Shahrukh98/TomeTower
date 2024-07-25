package dto

import (
	"tometower/internal/entity"
)

type CreateGenreDto struct {
	Name string `json:"name"`
}

type GenreDataDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func MapCreateGenreToGenre(genreDto CreateGenreDto) (entity.Genre, error) {
	var genre entity.Genre

	genre.Name = genreDto.Name
	return genre, nil
}

func MapGenreToGenreDataDto(genre entity.Genre) (GenreDataDto, error) {
	var genreDataDto GenreDataDto

	genreDataDto.Id = genre.Id
	genreDataDto.Name = genre.Name
	return genreDataDto, nil
}
