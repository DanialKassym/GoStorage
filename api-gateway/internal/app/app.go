package app

import (
	"log/slog"
	"net/http"
	"os"

	grpc_auth_client "github.com/DanialKassym/GoStorage/api-gateway/internal/client/auth_grpc_client"
	handlers "github.com/DanialKassym/GoStorage/api-gateway/internal/handlers"
	"github.com/rs/cors"
)

type App struct {
	HTTPserver *http.ServeMux
	Authclient *grpc_auth_client.GRPCClient
	logger     *slog.Logger
	HTTPAddr   string
}

func New(
	log *slog.Logger,
	grpcAddr string,
	httpAddr string,
) *App {
	grpcApp, err := grpc_auth_client.NewGRPCClient(grpcAddr)
	if err != nil {
		log.Error("Error initializing grpc client")
		os.Exit(1)
	}
	httpApp := http.NewServeMux()
	return &App{
		HTTPserver: httpApp,
		Authclient: grpcApp,
		logger:     log,
		HTTPAddr:   httpAddr,
	}
}

func (a *App) Run() {
	a.InitRoute()

	err := a.RunHTTPServer()
	if err != nil {
		a.logger.Error("couldnt init http server", err)
		os.Exit(1)
	}
}

func (a *App) InitRoute() {
	a.HTTPserver.HandleFunc("POST /login", handlers.Login(a.logger, a.Authclient))
	//a.HTTPserver.HandleFunc("POST /register", handlers.Register)
	//a.HTTPserver.HandleFunc("GET /validate-token", handlers.ValidateToken)
	//a.HTTPserver.HandleFunc("POST /upload/", handlers.UploadObject)
}

func (a *App) RunHTTPServer() error {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"Authorization"},
		Debug:            false,
	})
	a.logger.Warn(a.HTTPAddr)
	err := http.ListenAndServe(a.HTTPAddr, c.Handler(a.HTTPserver))
	if err != nil {
		return err
	}

	return nil
}
