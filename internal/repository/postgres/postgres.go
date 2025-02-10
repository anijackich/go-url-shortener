package postgres

import (
	"errors"
	"fmt"
	"github.com/anijackich/go-url-shortener/internal/models"
	"github.com/anijackich/go-url-shortener/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type LinkRepository struct {
	db *gorm.DB
}

func NewLinkRepository(
	DBHost string,
	DBPort int,
	DBUser string,
	DBPassword string,
	DBName string,
) (*LinkRepository, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBPassword, DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.AutoMigrate(&models.Link{}); err != nil {
		return nil, fmt.Errorf("failed to auto-migrate database: %w", err)
	}

	return &LinkRepository{db: db}, nil
}

func (r *LinkRepository) CreateLink(link *models.Link) error {
	if err := r.db.Create(link).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return repository.ErrLinkAlreadyExists
		}
		return err
	}

	return nil
}

func (r *LinkRepository) GetLinkByCode(code string) (*models.Link, error) {
	var link models.Link

	if err := r.db.Where("short_code = ?", code).First(&link).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrLinkNotFound
		}
		return nil, err
	}

	return &link, nil
}
