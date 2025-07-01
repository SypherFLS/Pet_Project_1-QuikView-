package main

import (
	"log/slog"
	"os"
	"project/internal/config"
	"project/internal/repository/postgres"
)

const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)

func setupLogger (env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}

func main() {
	cfg := config.MustLoad()

	db.InitDB()

	log := setupLogger(cfg.Env)

	log.Info("starting backend", slog.String("env", cfg.Env))


}


