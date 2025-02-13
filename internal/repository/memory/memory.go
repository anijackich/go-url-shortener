package memory

import (
	"github.com/anijackich/go-url-shortener/internal/models"
	"github.com/anijackich/go-url-shortener/internal/repository"
	"sync"
)

type LinkRepository struct {
	linksByCode map[string]*models.Link
	codesByUrl  map[string]string
	mu          sync.RWMutex
}

func NewLinkRepository() *LinkRepository {
	return &LinkRepository{
		linksByCode: make(map[string]*models.Link),
		codesByUrl:  make(map[string]string),
	}
}

func (r *LinkRepository) CreateLink(link *models.Link) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.linksByCode[link.ShortCode]; exists {
		return repository.ErrLinkAlreadyExists
	}

	if _, exists := r.codesByUrl[link.URL]; exists {
		return repository.ErrLinkAlreadyExists
	}

	r.linksByCode[link.ShortCode] = link
	r.codesByUrl[link.URL] = link.ShortCode

	return nil
}

func (r *LinkRepository) GetLinkByCode(code string) (*models.Link, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	link, exists := r.linksByCode[code]
	if !exists {
		return nil, repository.ErrLinkNotFound
	}

	return link, nil
}
