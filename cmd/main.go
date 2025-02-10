package main

import (
	"fmt"
	"go-url-shortener/internal/repository/memory"
	"go-url-shortener/internal/routers"
	"log"

	"github.com/gin-gonic/gin"
	"go-url-shortener/api/swagger"
	"go-url-shortener/internal/config"
	"go-url-shortener/internal/handlers"
	"go-url-shortener/internal/service"
)

// @title		URL Shortener API
// @version	0.1.0
// @BasePath	/api/v1
// @accept		json
// @produce	json
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

	linkHandler := handlers.NewLinkHandler(linkService)

	r := gin.Default()

	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	routers.SetupLinkRouter(v1, linkHandler)

	swagger.Setup(r)

	err = r.Run(addr)
	if err != nil {
		log.Fatalf("Failed to launch server: %s", err.Error())
		return
	}
}
