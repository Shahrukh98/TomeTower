package entity

import (
	"errors"
	"time"
)

type Role int

const (
	AdminRole Role = iota + 1
	UserRole
)

func RoleToString(s Role) (string, error) {
	roleMap := map[Role]string{
		AdminRole: "Admin",
		UserRole:  "User",
	}
	result, ok := roleMap[s]
	if !ok {
		return "", errors.New("bad role")
	}
	return result, nil
}

func StringToRole(str string) (Role, error) {
	roleStringToEnum := map[string]Role{
		"Admin": AdminRole,
		"User":  UserRole,
	}
	result, ok := roleStringToEnum[str]
	if !ok {
		return 0, errors.New("bad role")
	}
	return result, nil
}

type User struct {
	Id            string
	Name          string
	Nick          string
	Email         string
	Password      string
	Role          Role
	NickUpdatedAt time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

const NickUpdateCooldown = 3600
