package app

import (
	"log/slog"
	"os"

	grpc_server "github.com/DanialKassym/GoStorage/auth-service/internal/grpc_server"
)

type App struct {
	Authclient *grpc_server.GrpcServer
	JWTkey     string
	logger     *slog.Logger
}

func New(
	log *slog.Logger,
	grpcAddr string,
	jwtkey string,
) *App {
	grpcApp, err := grpc_server.NewGRPCServer(grpcAddr)
	if err != nil {
		log.Error("Error initializing grpc client")
		os.Exit(1)
	}
	return &App{
		Authclient: grpcApp,
		JWTkey: jwtkey,
		logger:     log,
	}
}

func (a *App) Run() {

}
