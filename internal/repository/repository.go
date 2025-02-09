package repository

import "go-url-shortener/internal/models"

type LinkRepository interface {
	CreateLink(link *models.Link) error
	GetLinkByCode(code string) (*models.Link, error)
}
