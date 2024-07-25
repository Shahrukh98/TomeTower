package service

import (
	"errors"

	"tometower/internal/dto"
	"tometower/internal/middleware"
	"tometower/pkg/util"
)

type AuthService struct {
	service UserService
}

func NewAuthService(service UserService) *AuthService {
	return &AuthService{service: service}
}

func (a *AuthService) Register(userDto dto.UserCreateDto) error {
	user, err := dto.MapCreateDtoToUser(&userDto)
	if err != nil {
		return err
	}

	return a.service.AddUser(*user)
}

func (a *AuthService) Login(email string, password string) (*dto.UserDataDto, error) {
	user, err := a.service.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	err = util.VerifyPassword(user.Password, password)
	if err != nil {
		return nil, errors.New("wrong pass")
	}

	token, err := middleware.CreateToken(user.Id, user.Name)
	if err != nil {
		return nil, errors.New("cant make token")
	}

	userData, err := dto.MapUserToUserDto(&user)
	if err != nil {
		return nil, err
	}
	userData.Token = token

	return userData, nil
}
