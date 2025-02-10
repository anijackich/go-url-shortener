package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/anijackich/go-url-shortener/internal/routers"
	"github.com/anijackich/go-url-shortener/internal/structs"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/anijackich/go-url-shortener/internal/handlers"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) ShortenLink(url string) (string, error) {
	args := m.Called(url)
	return args.String(0), args.Error(1)
}

func (m *MockService) ExpandShortLink(url string) (string, error) {
	args := m.Called(url)
	return args.String(0), args.Error(1)
}

func makeRequest(
	engine *gin.Engine,
	method string,
	baseUrl string,
	queryParams url.Values,
	body interface{},
) (int, *bytes.Buffer) {
	var reqBody []byte
	if body != nil {
		reqBody, _ = json.Marshal(body)
	} else {
		reqBody = nil
	}

	reqURL, _ := url.Parse(baseUrl)
	reqURL.RawQuery = queryParams.Encode()

	req, _ := http.NewRequest(
		method,
		reqURL.String(),
		bytes.NewBuffer(reqBody),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	engine.ServeHTTP(rec, req)

	return rec.Code, rec.Body
}

func TestShortenLink(t *testing.T) {
	testLongUrl := "https://some-site.com/long-url"
	testShortenedLink := "https://example.com/3bac1"

	mockService := &MockService{}
	linkHandler := handlers.NewLinkHandler(mockService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	routers.SetupLinkRouter(r.Group("/"), linkHandler)

	mockService.On(
		"ShortenLink",
		testLongUrl,
	).Return(testShortenedLink, nil)

	resCode, rawResBody := makeRequest(
		r,
		"POST",
		"/shorten",
		nil,
		structs.LongLink{URL: testLongUrl},
	)

	var resBody structs.ShortenedLink
	if err := json.NewDecoder(rawResBody).Decode(&resBody); err != nil {
		assert.NoError(t, err)
	}

	assert.Equal(t, http.StatusOK, resCode)
	assert.Equal(t, testShortenedLink, resBody.URL)

	mockService.AssertExpectations(t)
}

func TestExpandShortLink(t *testing.T) {
	testExpandedLink := "https://some-site.com/long-url"
	testShortUrl := "https://example.com/3bac1"

	mockService := &MockService{}
	linkHandler := handlers.NewLinkHandler(mockService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	routers.SetupLinkRouter(r.Group("/"), linkHandler)

	mockService.On(
		"ExpandShortLink",
		testShortUrl,
	).Return(testExpandedLink, nil)

	resCode, resRawBody := makeRequest(
		r,
		"GET",
		"/expand",
		url.Values{"u": []string{testShortUrl}},
		nil,
	)

	var resBody structs.LongLink
	if err := json.NewDecoder(resRawBody).Decode(&resBody); err != nil {
		assert.NoError(t, err)
	}

	assert.Equal(t, http.StatusOK, resCode)
	assert.Equal(t, testExpandedLink, resBody.URL)

	mockService.AssertExpectations(t)

}
