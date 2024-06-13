package user

import "time"

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
