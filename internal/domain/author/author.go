package author

import (
	"time"
)

type Author struct {
	ID          string
	Name        string
	PhotoUrl    string
	Nationality string
	DateOfBirth time.Time
	DateOfDeath time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
