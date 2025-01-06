package model

import "time"

type Author struct {
	ID        int
	Name      string
	Born      string
	Died      string
	Books     []*Book `gorm:"many2many:book_authors;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
