package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Host             string
	Port             int
	Domain           string
	LinkCodeLength   int
	LinkCodeAlphabet string
}

type DBConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

type Config struct {
	App AppConfig
	DB  DBConfig
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("No .env file found")
	}

	appPort, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}

	dbPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		return nil, err
	}

	linkCodeLength, err := strconv.Atoi(os.Getenv("LINK_CODE_LENGTH"))
	if err != nil {
		return nil, err
	}

	return &Config{
		App: AppConfig{
			Host:             os.Getenv("HOST"),
			Port:             appPort,
			Domain:           os.Getenv("DOMAIN"),
			LinkCodeLength:   linkCodeLength,
			LinkCodeAlphabet: os.Getenv("LINK_CODE_ALPHABET"),
		},
		DB: DBConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     dbPort,
			Name:     os.Getenv("POSTGRES_DATABASE"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
		},
	}, nil
}
