package model

import "time"

type Genre struct {
	ID        int
	Name      string `gorm:"UNIQUE"`
	Books     []*Book `gorm:"many2many:book_genres;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
