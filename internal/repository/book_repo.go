package repository

import (
	"context"

	"github.com/AssassinRobot/author/internal/model"
)

type BookRepository interface {
	GetAllBooks(ctx context.Context)([]*model.Book,error)
	GetBookByID(ctx context.Context, ID int) (*model.Book, error)
	GetBooksByLanguageID(ctx context.Context,languageID int) ([]*model.Book, error)
	GetBooksByGenreID(ctx context.Context,genreID int) ([]*model.Book, error)
	GetBooksByAuthorID(ctx context.Context,authorID int) ([]*model.Book, error)
	GetBooksByPublicationDate(ctx context.Context,publicationData int) ([]*model.Book, error)
	GetBooksByName(ctx context.Context, name string) ([]*model.Book, error)
	CreateBook(ctx context.Context, BookModel *model.Book) (*model.Book, error)
	UpdateBookByID(ctx context.Context, ID int, BookModel *model.Book) (*model.Book, error)
	DeleteBookByID(ctx context.Context, ID int) error
}
