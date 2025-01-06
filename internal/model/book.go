package model

import "time"

type Book struct {
	ID          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Publication int
	Authors     []*Author
	Genres      []*Genre
	Pages       int
	Language    *Language
}
