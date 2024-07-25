package repository

import "tometower/internal/entity"

type GenreRepository interface {
	Add(genre entity.Genre) (string, error)
	GetAll() ([]entity.Genre, error)
	GetById(id string) (entity.Genre, error)
	Remove(id string) error
}
