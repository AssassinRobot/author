package repo

import (
	"context"
	"errors"

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

	// err := a.db.WithContext(ctx).Preload("Books").Find(&languages).Error

	err := a.db.WithContext(ctx).Find(&languages).Error

	return languages, err
}

func (a *languageRepository) GetLanguageByID(ctx context.Context, ID int) (*model.Language, error) {
	language := new(model.Language)

	err := a.db.WithContext(ctx).First(language, ID).Error
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
	savedLanguage, err := a.GetLanguageByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	savedLanguage.Name = name

	return savedLanguage, nil
}

func (a *languageRepository) DeleteLanguageByID(ctx context.Context, ID int) error {
	language := new(model.Language)

	result := a.db.WithContext(ctx).Delete(language, ID)
	if result.Error != nil {
		return result.Error
	}

	// if result.RowsAffected == 0 {
	// 	return
	// }

	return nil
}
