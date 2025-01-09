package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/AssassinRobot/author/internal/model"
	"github.com/AssassinRobot/author/internal/repository"
	"gorm.io/gorm"
)

type languageRepository struct {
	db *gorm.DB
}

func NewLanguageRepository(db *gorm.DB) repository.LanguageRepository {
	return &languageRepository{
		db: db,
	}
}

func (a *languageRepository) GetAllLanguages(ctx context.Context) ([]*model.Language, error) {
	var languages = []*model.Language{}

	err := a.db.WithContext(ctx).Preload("Books").Preload("Books.Genres").Preload("Books.Language").Find(&languages).Error

	return languages, err
}

func (a *languageRepository) GetLanguageByID(ctx context.Context, ID int) (*model.Language, error) {
	language := new(model.Language)

	err := a.db.WithContext(ctx).Preload("Books").Preload("Books.Genres").Preload("Books.Language").First(language, ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLanguageNotFound
		}
		return nil, err
	}

	return language, nil
}

func (a *languageRepository) CreateLanguage(ctx context.Context, languageModel *model.Language) (*model.Language, error) {
	tx := a.db.Create(languageModel)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return languageModel, nil
}

func (a *languageRepository) UpdateLanguageByID(ctx context.Context, ID int, name string) (*model.Language, error) {
	err := a.db.Model(&model.Language{}).Where("id = ?", ID).Update("name", name).Error
	if err != nil {
		return nil, err
	}

	language,err := a.GetLanguageByID(ctx,ID)
	if err != nil {
		return nil, err
	}
	
	return language, nil
}

func (a *languageRepository) DeleteLanguageByID(ctx context.Context, ID int) error {
	var count int64
	err := a.db.WithContext(ctx).Model(&model.Book{}).Where("language_id = ?", ID).Count(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("cannot delete language: there are %d books associated with this language", count)
	}

	result := a.db.WithContext(ctx).Delete(&model.Language{}, ID)
	if result.Error != nil {
		return result.Error
	}

	// if result.RowsAffected == 0 {
	// 	return
	// }

	return nil
}
