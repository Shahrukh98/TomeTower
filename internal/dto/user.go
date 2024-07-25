package dto

import (
	"tometower/internal/entity"
)

type UserCreateDto struct {
	Name     string `json:"name"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserLoginDto struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

type UserDataDto struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Nick  string `json:"nick"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

type NickUpdateDto struct {
	Nick string
}

func MapUserToUserDto(user *entity.User) (*UserDataDto, error) {
	var userData UserDataDto

	userData.Id = user.Id
	userData.Name = user.Name
	userData.Nick = user.Nick
	userData.Email = user.Email
	role, err := entity.RoleToString(user.Role)
	if err != nil {
		return nil, err
	}
	userData.Role = role
	return &userData, nil
}

func MapCreateDtoToUser(userCreateDto *UserCreateDto) (*entity.User, error) {
	var user entity.User
	role, err := entity.StringToRole(userCreateDto.Role)
	if err != nil {
		return nil, err
	}

	user.Name = userCreateDto.Name
	user.Nick = userCreateDto.Nick
	user.Email = userCreateDto.Email
	user.Password = userCreateDto.Password
	user.Role = role

	return &user, nil
}
