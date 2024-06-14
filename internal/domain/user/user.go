package user

import (
	"time"
)

type User struct {
	ID            string
	Name          string
	Nick          string
	Email         string
	Password      string
	NickUpdatedAt time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type NickUpdate struct {
	Nick string
}

const NickUpdateCooldown = 3600

type UserLoginResponse struct {
	ID    string
	Name  string
	Nick  string
	Token string
}

type UserLoginRequest struct {
	Email    string
	Password string
}
