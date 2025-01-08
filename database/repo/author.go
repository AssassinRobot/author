package repo

import (
	"context"
	"errors"

	"github.com/AssassinRobot/author/internal/model"
	"github.com/AssassinRobot/author/internal/repository"
	"gorm.io/gorm"
)

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) repository.AuthorRepository {
	return &authorRepository{
		db: db,
	}
}

func (r *authorRepository) FindByIDs(ctx context.Context, IDs []string) ([]*model.Author, error) {
	var authors []*model.Author
	err := r.db.WithContext(ctx).Where("id IN ?", IDs).Find(&authors).Error
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func (a *authorRepository) GetAllAuthors(ctx context.Context) ([]*model.Author, error) {
	var authors = []*model.Author{}

	err := a.db.WithContext(ctx).Preload("Books").Find(&authors).Error

	return authors, err
}

func (a *authorRepository) GetAuthorByID(ctx context.Context, ID int) (*model.Author, error) {
	Author := new(model.Author)

	err := a.db.WithContext(ctx).Preload("Books").First(Author, ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAuthorNotFound
		}
		return nil, err
	}

	return Author, nil
}

func (a *authorRepository) GetAuthorsByNames(ctx context.Context, name string) ([]*model.Author, error) {
	var authors []*model.Author

	err := a.db.WithContext(ctx).Preload("Books").Where("name = ?", name).Find(&authors).Error

	if err != nil {
		return nil, err
	}
	// if len(authors) == 0 {
	// 		return nil, ErrBookNotFound
	// }

	return authors, err
}

func (a *authorRepository) CreateAuthor(ctx context.Context, authorModel *model.Author) (*model.Author, error) {
	tx := a.db.Create(authorModel)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return authorModel, nil
}

func (a *authorRepository) UpdateAuthorByID(ctx context.Context, ID int, authorModel *model.Author) (*model.Author, error) {
	savedAuthor, err := a.GetAuthorByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	savedAuthor.Name = authorModel.Name
	savedAuthor.Born = authorModel.Born
	savedAuthor.Died = authorModel.Died

	err = a.db.Model(savedAuthor).Association("Books").Replace(authorModel.Books)
	if err != nil {
		return nil, err
	}

	return savedAuthor, nil
}

func (a *authorRepository) DeleteAuthorByID(ctx context.Context, ID int) error {
	Author := new(model.Author)

	result := a.db.WithContext(ctx).Delete(Author, ID)
	if result.Error != nil {
		return result.Error
	}

	// if result.RowsAffected == 0 {
	// 	return
	// }

	return nil
}
