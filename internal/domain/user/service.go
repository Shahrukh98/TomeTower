package user

import (
	"tometower/pkg/util"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) AddUser(user User) error {
	hashed_password, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed_password
	return s.repo.Add(user)
}

func (s *UserService) FindUserById(id string) (User, error) {
	return s.repo.FindByID(id)
}
