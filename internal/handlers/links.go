package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-url-shortener/internal/repository"
	"go-url-shortener/internal/service"
	"go-url-shortener/internal/structs"
	"net/http"
)

type LinkServiceInterface interface {
	ShortenLink(url string) (string, error)
	ExpandShortLink(url string) (string, error)
}

type LinkHandler struct {
	service LinkServiceInterface
}

func NewLinkHandler(service LinkServiceInterface) *LinkHandler {
	return &LinkHandler{service: service}
}

// ShortenLink
//
//	@Summary		Shorten link
//	@Description	Converts a long URL to a short link
//	@Accept			json
//	@Produce		json
//	@Param			long_url	body	structs.LongLink	true	"Long URL"
//	@Router			/shorten [post]
func (h *LinkHandler) ShortenLink(c *gin.Context) {
	var longLink structs.LongLink

	err := c.BindJSON(&longLink)
	if err != nil {
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{"error": err},
		)
	}

	shortenedLink, err := h.service.ShortenLink(longLink.URL)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrLinkAlreadyExists):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errors.Is(err, service.ErrInvalidURL):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(
		http.StatusOK,
		structs.ShortenedLink{
			URL: shortenedLink,
		},
	)
}

// ExpandLink
//
//	@Summary		Expand link
//	@Description	Returns a long URL for the specified link
//	@Accept			json
//	@Produce		json
//	@Param			u	query	string	true	"Short link"
//	@Router			/expand [get]
func (h *LinkHandler) ExpandLink(c *gin.Context) {
	shortLink := c.Query("u")

	expandedLink, err := h.service.ExpandShortLink(shortLink)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrLinkNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case errors.Is(err, service.ErrInvalidURL):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(
		http.StatusOK,
		structs.LongLink{
			URL: expandedLink,
		},
	)
}
