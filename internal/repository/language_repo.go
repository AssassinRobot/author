package repository

import (
	"context"

	"github.com/AssassinRobot/author/internal/model"
)

type LanguageRepository interface {
	GetAllLanguages(ctx context.Context) ([]*model.Language, error)
	CreateLanguage(ctx context.Context, languageModel *model.Language) (*model.Language, error)
	UpdateLanguageByID(ctx context.Context, ID int, name string) (*model.Language, error)
	DeleteLanguageByID(ctx context.Context, ID int) error
}
