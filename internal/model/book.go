package model

import "time"

type Book struct {
	ID          int
	Name        string
	Publication int
	Pages       int
	LanguageID  int      
	Language    *Language `gorm:"foreignKey:LanguageID"` 
	Authors     []*Author `gorm:"many2many:book_authors;"`
	Genres      []*Genre  `gorm:"many2many:book_genres;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

