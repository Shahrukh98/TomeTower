package repository

import "tometower/internal/entity"

type AuthorRepository interface {
	Add(author entity.Author) (string, error)
	GetAll() ([]entity.Author, error)
	GetById(id string) (entity.Author, error)
	Remove(id string) error
}
