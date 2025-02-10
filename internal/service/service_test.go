package service_test

import (
	"fmt"
	"github.com/anijackich/go-url-shortener/internal/models"
	"github.com/anijackich/go-url-shortener/internal/repository"
	"github.com/anijackich/go-url-shortener/internal/service"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateLink(link *models.Link) error {
	args := m.Called(link)
	return args.Error(0)
}

func (m *MockRepository) GetLinkByCode(code string) (*models.Link, error) {
	args := m.Called(code)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*models.Link), args.Error(1)
}

func TestNewLinkServiceInvalidDomain(t *testing.T) {
	testCodeAlphabet := "abc123"
	testCodeLength := 5

	invalidCases := []string{
		".",
		"hello",
		"https://hello.com",
	}

	repo := &MockRepository{}
	for _, invalidDomain := range invalidCases {
		_, err := service.NewLinkService(
			invalidDomain,
			testCodeAlphabet,
			testCodeLength,
			repo,
		)

		if assert.Error(t, err) {
			assert.Equal(t, service.ErrInvalidDomain, err)
		}
	}
}

func TestShortenLink(t *testing.T) {
	testDomain := "example.com"
	testCodeAlphabet := "abc123"
	testCodeLength := 5
	testLongUrl := "https://some-site.com/long-url"

	repo := &MockRepository{}
	linkService, err := service.NewLinkService(
		testDomain,
		testCodeAlphabet,
		testCodeLength,
		repo,
	)
	assert.NoError(t, err)

	repo.On(
		"CreateLink",
		mock.AnythingOfType("*models.Link"),
	).Return(nil)

	shortenedLink, err := linkService.ShortenLink(testLongUrl)
	assert.NoError(t, err)

	assert.Regexpf(
		t,
		fmt.Sprintf(`^https://%s/[^/]+$`, testDomain),
		shortenedLink,
		"\"%s\" does not satisfy valid URL with \"%s\" domain",
		shortenedLink,
		testDomain,
	)

	repo.AssertExpectations(t)
}

func TestShortenLinkInvalidURL(t *testing.T) {
	testDomain := "example.com"
	testCodeAlphabet := "abc123"
	testCodeLength := 5

	invalidCases := []string{
		"",
		"http",
		"https://",
		"https://www",
		"abcde12345!@#$%",
	}

	repo := &MockRepository{}
	linkService, err := service.NewLinkService(
		testDomain,
		testCodeAlphabet,
		testCodeLength,
		repo,
	)
	assert.NoError(t, err)

	for _, invalidURL := range invalidCases {
		_, err = linkService.ShortenLink(invalidURL)

		if assert.Error(t, err) {
			assert.Equal(t, service.ErrInvalidURL, err)
		}
	}

}

func TestExpandShortLink(t *testing.T) {
	testDomain := "example.com"
	testCodeAlphabet := "abc123"
	testCodeLength := 5
	testLongUrl := "https://some-site.com/long-url"
	testShortCode := "3bac1"

	repo := &MockRepository{}
	linkService, err := service.NewLinkService(
		testDomain,
		testCodeAlphabet,
		testCodeLength,
		repo,
	)
	assert.NoError(t, err)

	repo.On("GetLinkByCode", testShortCode).Return(
		&models.Link{
			URL:       testLongUrl,
			ShortCode: testShortCode,
		}, nil,
	)

	shortUrl := url.URL{
		Scheme: "https",
		Host:   testDomain,
		Path:   testShortCode,
	}

	expandedLink, err := linkService.ExpandShortLink(shortUrl.String())
	assert.NoError(t, err)

	assert.Equal(t, testLongUrl, expandedLink)

	repo.AssertExpectations(t)
}

func TestExpandShortLinkNotFound(t *testing.T) {
	testDomain := "example.com"
	testCodeAlphabet := "abc123"
	testCodeLength := 5
	testShortCode := "3bac1"

	repo := &MockRepository{}
	linkService, err := service.NewLinkService(
		testDomain,
		testCodeAlphabet,
		testCodeLength,
		repo,
	)
	assert.NoError(t, err)

	repo.On("GetLinkByCode", testShortCode).Return(
		nil,
		repository.ErrLinkNotFound,
	)

	shortUrl := url.URL{
		Scheme: "https",
		Host:   testDomain,
		Path:   testShortCode,
	}

	_, err = linkService.ExpandShortLink(shortUrl.String())

	if assert.Error(t, err) {
		assert.Equal(t, repository.ErrLinkNotFound, err)
	}

	repo.AssertExpectations(t)
}
