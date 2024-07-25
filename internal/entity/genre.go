package entity

import (
	"time"
)

type Genre struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
