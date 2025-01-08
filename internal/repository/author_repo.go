package repository

import (
	"context"

	"github.com/AssassinRobot/author/internal/model"
)

type AuthorRepository interface {
	GetAllAuthors(ctx context.Context) ([]*model.Author, error)
	FindByIDs(ctx context.Context, IDs []string) ([]*model.Author, error)
	GetAuthorByID(ctx context.Context, ID int) (*model.Author, error)
	GetAuthorsByNames(ctx context.Context, name string) ([]*model.Author, error)
	CreateAuthor(ctx context.Context, authorModel *model.Author) (*model.Author, error)
	UpdateAuthorByID(ctx context.Context, ID int, authorModel *model.Author) (*model.Author, error)
	DeleteAuthorByID(ctx context.Context, ID int) error
}
