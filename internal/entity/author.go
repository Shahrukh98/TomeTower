package entity

import (
	"time"
)

type Author struct {
	Id          string
	Name        string
	PhotoUrl    string
	Nationality string
	DateOfBirth string
	DateOfDeath string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
