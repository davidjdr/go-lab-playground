package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()

	r := gin.New()
	// Agrega un middleware que recupera de cualquier pánico y evita que el servidor se caiga
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })

	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	logger.Info("starting api", zap.String("addr", addr))
	if err := r.Run(addr); err != nil {
		logger.Fatal("server error", zap.Error(err))
	}
}
