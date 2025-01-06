package model

import "time"

type Language struct {
	ID        int
	Name      string
	Books     []*Book `gorm:"foreignKey:LanguageID"` 
	CreatedAt time.Time
	UpdatedAt time.Time
}
