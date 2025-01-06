package model

import "time"

type Genre struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
