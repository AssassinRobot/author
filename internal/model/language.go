package model

import "time"

type Language struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
