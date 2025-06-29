package app

import (
	"log/slog"
	"net/http"
	"github.com/DanialKassym/GoStorage/api-gateway/internal/handler"
)

type App struct {
	HTTPserver *http.ServeMux
	Authclient *auth_client.App
}

func New(
	log *slog.Logger,
	grpcAddr string,
) *App {
	grpcApp := auth_client.NewGrpcClient(log, grpcAddr)
	httpApp := http.NewServeMux()
	handler.InitRoute(httpApp)
	return &App{
		HTTPserver: httpApp,
		Authclient: grpcApp,
	}
}

func (a *App) Run(
	log *slog.Logger,
	httpAddr string,
	server *http.ServeMux,
) {
	err := handler.RunHTTPServer(server, httpAddr)
	if err != nil {
		log.Warn("couldnt init http server")
	}
}
