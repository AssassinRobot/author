package repo

import (
	"context"
	"errors"

	"github.com/AssassinRobot/author/internal/model"
	"github.com/AssassinRobot/author/internal/repository"
	"gorm.io/gorm"
)

type genreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) repository.GenreRepository {
	return &genreRepository{
		db: db,
	}
}

func (r *genreRepository) FindByIDs(ctx context.Context, IDs []int) ([]*model.Genre, error) {
	var genres []*model.Genre
	err := r.db.WithContext(ctx).Where("id IN ?", IDs).Find(&genres).Error
	if err != nil {
		return nil, err
	}
	
	return genres, nil
}

func (a *genreRepository) GetAllGenres(ctx context.Context) ([]*model.Genre, error) {
	var genres = []*model.Genre{}

	err := a.db.WithContext(ctx).Preload("Books").Find(&genres).Error

	// err := a.db.WithContext(ctx).Find(&genres).Error

	return genres, err
}

func (a *genreRepository) GetGenreByID(ctx context.Context, ID int) (*model.Genre, error) {
	genre := new(model.Genre)

	err := a.db.WithContext(ctx).First(genre, ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGenreNotFound
		}
		return nil, err
	}

	return genre, nil
}

// func (a *genreRepository) GetAuthorsByNames(ctx context.Context, name string) ([]*model.genre, error) {
// 	var genres []*model.Genre

// 	err := a.db.WithContext(ctx).Preload("Books").Where("name = ?", name).Find(&genres).Error

// 	if err != nil {
// 		return nil, err
// 	}
// 	// if len(genres) == 0 {
// 	// 		return nil, ErrBookNotFound
// 	// }

// 	return genres, err
// }

func (a *genreRepository) CreateGenre(ctx context.Context, genreModel *model.Genre) (*model.Genre, error) {
	tx := a.db.Create(genreModel)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return genreModel, nil
}

func (a *genreRepository) UpdateGenreByID(ctx context.Context, ID int, name string) (*model.Genre, error) {
	savedGenre, err := a.GetGenreByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	savedGenre.Name = name

	return savedGenre, nil
}

func (a *genreRepository) DeleteGenreByID(ctx context.Context, ID int) error {
	genre := new(model.Genre)

	result := a.db.WithContext(ctx).Delete(genre, ID)
	if result.Error != nil {
		return result.Error
	}

	// if result.RowsAffected == 0 {
	// 	return
	// }

	return nil
}
