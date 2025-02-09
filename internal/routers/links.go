package routers

import (
	"github.com/gin-gonic/gin"
	"go-url-shortener/internal/handlers"
)

func SetupLinkRouter(
	group *gin.RouterGroup,
	linkHandler *handlers.LinkHandler,
) {
	group.POST("/shorten", linkHandler.ShortenLink)
	group.GET("/expand", linkHandler.ExpandLink)
}
