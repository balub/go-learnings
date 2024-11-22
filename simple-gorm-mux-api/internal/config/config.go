package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	File string
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file. Using environment variables.")
	}
}

func GetDBConfig() DBConfig {
	return DBConfig{
		File: os.Getenv("FILENAME"),
	}
}
