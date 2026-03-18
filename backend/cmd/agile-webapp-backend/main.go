package main

import (
	"log/slog"
	"os"

	"github.com/Kaiman30/AgileWebApp/backend/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setup_logger(cfg.Env)

	log.Info("starting application", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// TODO: init storage

	// TODO: init router

	// TODO: run server

	log.Info("server is running", slog.String("address", cfg.HTTPServer.Address))

}

func setup_logger(env string) *slog.Logger {
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
