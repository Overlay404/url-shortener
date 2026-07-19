package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/handlers"
	"url-shortener/internal/repositories"
	"url-shortener/internal/services"

	"github.com/gin-gonic/gin"
)

const (
	envLocal       = "local"
	envDevelopment = "dev"
	envProduction  = "prod"
)

func main() {
	conf, errConf := config.Load()
	if errConf != nil {
		log.Fatalf("Failed to load config %w", errConf)
	}

	log := slogLogger(conf.App.LevelLogs)
	log.Debug("Config loaded", "config",
		slog.String("App", conf.App.LevelLogs),
	)

	ctx := context.Background()
	repos := repositories.New(&ctx, conf, log)
	serv := services.New(repos, log)
	h := handlers.New(serv, log)

	router := gin.Default()
	router.GET("v1/g/:url", h.Get)
	router.POST("v1/set", h.Set)
	router.Run(":8000")
}

func slogLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProduction:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case envDevelopment:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	return log
}
