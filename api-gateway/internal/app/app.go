package app

import (
	"log/slog"
	"net/http"
	"os"

	grpc_auth_client "github.com/DanialKassym/GoStorage/api-gateway/internal/client/auth_grpc_client"
	handler"github.com/DanialKassym/GoStorage/api-gateway/internal/handler"
)

type App struct {
	HTTPserver *http.ServeMux
	Authclient *grpc_auth_client.GRPCClient
}

func New(
	log *slog.Logger,
	grpcAddr string,
) *App {
	grpcApp, err := grpc_auth_client.NewGRPCClient(grpcAddr)
	if err != nil {
		log.Error("Error initializing grpc client")
		os.Exit(1)
	}
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
