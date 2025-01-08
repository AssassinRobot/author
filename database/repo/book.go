package repo

import (
	"context"
	"errors"

	"github.com/AssassinRobot/author/internal/model"
	"github.com/AssassinRobot/author/internal/repository"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) repository.BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) FindByIDs(ctx context.Context, IDs []string)([]*model.Book, error) {
	var books []*model.Book
	err := r.db.WithContext(ctx).Where("id IN ?", IDs).Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *bookRepository) GetAllBooks(ctx context.Context) ([]*model.Book, error) {
	var books []*model.Book

	err := r.db.WithContext(ctx).Preload("Language").Preload("Authors").Preload("Genres").Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *bookRepository) GetBookByID(ctx context.Context, ID int) (*model.Book, error) {
	var book model.Book

	err := r.db.WithContext(ctx).Preload("Language").Preload("Authors").Preload("Genres").First(&book, ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBookNotFound
		}
		return nil, err
	}

	return &book, nil
}

func (r *bookRepository) GetBooksByLanguageID(ctx context.Context, languageID int) ([]*model.Book, error) {
	var books []*model.Book

	err := r.db.WithContext(ctx).Preload("Language").Preload("Authors").Preload("Genres").Where("language_id = ?", languageID).Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *bookRepository) GetBooksByGenreID(ctx context.Context, genreID int) ([]*model.Book, error) {
	var books []*model.Book

	err := r.db.WithContext(ctx).Preload("Language").Preload("Authors").Preload("Genres").Joins("JOIN book_genres ON book_genres.book_id = books.id").Where("book_genres.genre_id = ?", genreID).Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *bookRepository) GetBooksByAuthorID(ctx context.Context, authorID int) ([]*model.Book, error) {
	var books []*model.Book

	err := r.db.WithContext(ctx).Preload("Language").Preload("Authors").Preload("Genres").Joins("JOIN book_authors ON book_authors.book_id = books.id").Where("book_authors.author_id = ?", authorID).Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *bookRepository) GetBooksByPublicationDate(ctx context.Context, publicationDate int) ([]*model.Book, error) {
	var books []*model.Book
	err := r.db.WithContext(ctx).Preload("Language").Preload("Authors").Preload("Genres").Where("publication = ?", publicationDate).Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *bookRepository) GetBooksByName(ctx context.Context, name string) ([]*model.Book, error) {
	var books []*model.Book
	err := r.db.WithContext(ctx).Preload("Language").Preload("Authors").Preload("Genres").Where("name LIKE ?", "%"+name+"%").Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *bookRepository) CreateBook(ctx context.Context, bookModel *model.Book) (*model.Book, error) {
	err := r.db.WithContext(ctx).Create(bookModel).Error
	if err != nil {
		return nil, err
	}

	return bookModel, nil
}

func (r *bookRepository) UpdateBookByID(ctx context.Context, ID int, bookModel *model.Book) (*model.Book, error) {
	err := r.db.WithContext(ctx).Model(&model.Book{}).Where("id = ?", ID).Updates(bookModel).Error

	if err != nil {
		return nil, err
	}

	return bookModel, nil
}

func (r *bookRepository) DeleteBookByID(ctx context.Context, ID int) error {
	err := r.db.WithContext(ctx).Delete(&model.Book{}, ID).Error

	if err != nil {
		return err
	}
	return nil
}
