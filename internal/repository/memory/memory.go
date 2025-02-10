package memory

import (
	"github.com/anijackich/go-url-shortener/internal/models"
	"github.com/anijackich/go-url-shortener/internal/repository"
	"sync"
)

type LinkRepository struct {
	links map[string]*models.Link
	mu    sync.RWMutex
}

func NewLinkRepository() *LinkRepository {
	return &LinkRepository{
		links: make(map[string]*models.Link),
	}
}

func (r *LinkRepository) CreateLink(link *models.Link) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.links[link.ShortCode]; exists {
		return repository.ErrLinkAlreadyExists
	}

	r.links[link.ShortCode] = link

	return nil
}

func (r *LinkRepository) GetLinkByCode(code string) (*models.Link, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	link, exists := r.links[code]
	if !exists {
		return nil, repository.ErrLinkNotFound
	}

	return link, nil
}
