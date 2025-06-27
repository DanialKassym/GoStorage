package app

import (
	"log/slog"
	"net/http"

)

type App struct {
	HTTPSserver *http.ServeMux
	authClient  *auth_client.App
}

func New(
	log *slog.Logger,
	grpcPort int,
) *App {
	grpcApp := auth_client.NewGrpcClient(log, grpcPort)

	return &App{
		authClient: grpcApp,
	}
}
