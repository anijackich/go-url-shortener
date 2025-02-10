package service

import (
	"go-url-shortener/internal/models"
	"go-url-shortener/internal/repository"
	"go-url-shortener/pkg/utils"
	neturl "net/url"
	"regexp"
	"strings"
)

func isValidDomain(domain string) bool {
	re := regexp.MustCompile(
		`^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)(?:\.[a-z]{2,})+$`,
	)

	return re.MatchString(domain)
}

func isValidURL(url string) bool {
	re := regexp.MustCompile(
		`^https?://(?:www\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_+.~#?&/=]*)$`,
	)

	return re.MatchString(url)

}

type LinkService struct {
	domain       string
	codeAlphabet string
	codeLength   int
	repo         repository.LinkRepository
}

func NewLinkService(
	domain string,
	codeAlphabet string,
	codeLength int,
	linksRepo repository.LinkRepository,
) (*LinkService, error) {
	if !isValidDomain(domain) {
		return nil, ErrInvalidDomain
	}

	return &LinkService{
		domain:       domain,
		codeAlphabet: codeAlphabet,
		codeLength:   codeLength,
		repo:         linksRepo,
	}, nil
}

func (s *LinkService) ShortenLink(url string) (string, error) {
	if !isValidURL(url) {
		return "", ErrInvalidURL
	}

	parsedUrl, err := neturl.ParseRequestURI(url)
	if err != nil {
		return "", ErrInvalidURL
	}

	shortCode, err := utils.GenerateRandomString(s.codeAlphabet, s.codeLength)
	if err != nil {
		return "", err
	}

	err = s.repo.CreateLink(&models.Link{
		ShortCode: shortCode,
		URL:       parsedUrl.String(),
	})
	if err != nil {
		return "", err
	}

	shortUrl := neturl.URL{
		Scheme: "https",
		Host:   s.domain,
		Path:   shortCode,
	}

	return shortUrl.String(), nil
}

func (s *LinkService) ExpandShortLink(url string) (string, error) {
	parsedUrl, err := neturl.ParseRequestURI(url)
	if err != nil {
		return "", ErrInvalidURL
	}

	splitPath := strings.Split(parsedUrl.Path, "/")

	if len(splitPath) < 2 || len(splitPath) > 3 {
		return "", ErrInvalidURL
	}

	shortCode := strings.Split(parsedUrl.Path, "/")[1]

	link, err := s.repo.GetLinkByCode(shortCode)
	if err != nil {
		return "", err
	}

	return link.URL, nil
}
