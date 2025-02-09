package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go-url-shortener/internal/config"
	"go-url-shortener/internal/handlers"
	"go-url-shortener/internal/repository/memory"
	"go-url-shortener/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot load config: %s", err.Error())
		return
	}

	addr := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)

	linkRepo := memory.NewLinkRepository()

	linkService, err := service.NewLinkService(
		cfg.App.Domain,
		cfg.App.LinkCodeAlphabet,
		cfg.App.LinkCodeLength,
		linkRepo,
	)
	if err != nil {
		log.Fatalf("Cannot init link service: %s", err.Error())
		return
	}

	linkHandler := handlers.NewLinkHandler(*linkService)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/shorten", linkHandler.ShortenLink)
		v1.GET("/expand", linkHandler.ExpandLink)
	}

	err = r.Run(addr)
	if err != nil {
		log.Fatalf("Failed to launch server: %s", err.Error())
		return
	}
}
