package main

import (
	"log/slog"
	"os"

	"github.com/DanialKassym/GoStorage/auth-service/internal/app"
	"github.com/DanialKassym/GoStorage/auth-service/internal/config"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		slog.Error("Couldnt load .env file")
		os.Exit(1)
	}
	cfg := config.NewConfig()
	log := setupLogger()

	application := app.New(log, cfg.GRPCAddr,cfg.JWTKey)
	application.Run()
}

func setupLogger() *slog.Logger {
	var log *slog.Logger
	log = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	return log
}
