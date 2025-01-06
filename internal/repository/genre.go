package repository

import (
	"context"

	"github.com/AssassinRobot/author/internal/model"
)

type GenreRepository interface {
	GetAllGenres(ctx context.Context) ([]*model.Genre, error)
	CreateGenre(ctx context.Context, genreModel *model.Genre) (*model.Genre, error)
	UpdateGenreByID(ctx context.Context, ID int, name string) (*model.Genre, error)
	DeleteGenreByID(ctx context.Context, ID int) error
}
