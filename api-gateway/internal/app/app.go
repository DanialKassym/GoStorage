package app

import (
	"log/slog"
	"net/http"

	"github.com/DanialKassym/GoStorage/api-gateway/internal/client/auth_grpc_client"
	"github.com/DanialKassym/GoStorage/api-gateway/internal/handler"
)

type App struct {
	HTTPserver *http.ServeMux
	Authclient *auth_grpc_client.GRPCClient
}

func New(
	log *slog.Logger,
	grpcAddr string,
) *App {
	grpcApp := 
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
