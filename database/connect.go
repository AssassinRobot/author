package database

import (
	"sync"

	"github.com/AssassinRobot/author/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	postgresInstance *gorm.DB
	mu               = &sync.Mutex{}
)

func GetPostgresqlDB(dbURL string) (*gorm.DB, error) {
	mu.Lock()
	defer mu.Unlock()

	if postgresInstance == nil {
		db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		postgresInstance = db
	}

	migration(model.Author{})
	migration(model.Book{})
	migration(model.Genre{})
	migration(model.Language{})

	return postgresInstance, nil
}

func migration(model any) {
	err := postgresInstance.AutoMigrate(&model)
	if err != nil {
		panic(err)
	}
}
