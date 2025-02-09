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
	Port     string
	Name     string
	User     string
	Password string
}

type Config struct {
	App AppConfig
	DB  DBConfig
}

func toInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Cannot convert %s to int\n", s)
	}

	return num
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("No .env file found")
	}

	return &Config{
		App: AppConfig{
			Host:             os.Getenv("HOST"),
			Port:             toInt(os.Getenv("PORT")),
			Domain:           os.Getenv("DOMAIN"),
			LinkCodeLength:   toInt(os.Getenv("LINK_CODE_LENGTH")),
			LinkCodeAlphabet: os.Getenv("LINK_CODE_ALPHABET"),
		},
		DB: DBConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			Name:     os.Getenv("POSTGRES_DATABASE"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
		},
	}
}
