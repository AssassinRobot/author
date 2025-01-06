package model

import "time"

type Author struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Born      string
	Died      string
	Books     []*Book
}
