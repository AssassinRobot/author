package main

import (
	"log"

	"github.com/AssassinRobot/author/config"
	"github.com/AssassinRobot/author/database"
	"github.com/AssassinRobot/author/database/repo"
	"github.com/AssassinRobot/author/server"
)

func main() {
	serverPort, DBUrl := config.GetConfigs()
	db, err := database.GetPostgresqlDB(DBUrl)
	if err != nil {
		log.Fatal(err)
	}

	authorRepo := repo.NewAuthorRepository(db)
	bookRepo := repo.NewBookRepository(db)
	languageRepo := repo.NewLanguageRepository(db)
	genreRepo := repo.NewGenreRepository(db)

	server.InitializeServer(serverPort,authorRepo,bookRepo,genreRepo,languageRepo)
}
