package service

import (
	"errors"
	"time"

	"tometower/internal/entity"
	"tometower/internal/repository"
	"tometower/pkg/util"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) AddUser(user entity.User) error {
	hashed_password, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed_password
	return s.repo.Add(user)
}

func (s *UserService) GetByEmail(email string) (entity.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) GetUserById(id string) (entity.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) UpdateNick(id string, nick string) error {
	user, err := s.repo.GetById(id)
	if err != nil {
		return err
	}

	currentTime := time.Now().Unix()

	timeDiff := currentTime - user.NickUpdatedAt.Unix()
	if timeDiff < entity.NickUpdateCooldown {
		return errors.New("nick update on cooldown")
	}
	s.repo.UpdateNick(id, nick)

	return nil
}
