package routers

import (
	"github.com/anijackich/go-url-shortener/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupLinkRouter(
	group *gin.RouterGroup,
	linkHandler *handlers.LinkHandler,
) {
	group.POST("/shorten", linkHandler.ShortenLink)
	group.GET("/expand", linkHandler.ExpandLink)
}
