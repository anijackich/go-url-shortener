package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-url-shortener/internal/repository"
	"go-url-shortener/internal/service"
	"net/http"
)

type LongLink struct {
	URL string `json:"long_url"`
}

type ShortenedLink struct {
	URL string `json:"short_url"`
}

type LinkHandler struct {
	service service.LinkService
}

func NewLinkHandler(service service.LinkService) *LinkHandler {
	return &LinkHandler{service: service}
}

func (h *LinkHandler) ShortenLink(c *gin.Context) {
	var longLink LongLink

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
		ShortenedLink{
			URL: shortenedLink,
		},
	)
}

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
		LongLink{
			URL: expandedLink,
		},
	)
}
