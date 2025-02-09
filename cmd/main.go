package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-url-shortener/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	addr := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)

	r := gin.Default()

	r.Run(addr)
}
