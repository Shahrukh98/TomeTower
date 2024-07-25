package repository

import (
	"tometower/internal/entity"
)

type UserRepository interface {
	Add(user entity.User) error
	GetByEmail(email string) (entity.User, error)
	GetById(id string) (entity.User, error)
	UpdateNick(id string, nick string) error
}
