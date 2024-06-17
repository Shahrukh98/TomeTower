package genre

import (
	"time"
)

type Genre struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
