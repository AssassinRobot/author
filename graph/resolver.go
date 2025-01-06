package graph

import "github.com/AssassinRobot/author/internal/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthorRepo   repository.AuthorRepository
	BookRepo     repository.BookRepository
	LanguageRepo repository.LanguageRepository
	GenreRepo    repository.GenreRepository
}
