package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func configsDirPath() string {
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		panic("Error in generating env dir")
	}

	return filepath.Dir(f)
}

func GetConfigs() (string, string) {
	err := godotenv.Load(configsDirPath() + "/../" + "dev.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverPort := os.Getenv("SERVER_PORT")
	DBUrl := os.Getenv("DB_URL")

	return serverPort, DBUrl
}
