package handlers

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

// InitRoutes initializes all the routes for the application.
func InitRoute(server *http.ServeMux) {
	server.HandleFunc("POST /login", handlers.Login)
	server.HandleFunc("GET /get-access-token", handlers.Login)
	server.HandleFunc("POST /register", handlers.Register)
	server.HandleFunc("GET /validate-token", handlers.Main)
	server.HandleFunc("POST /upload/", handlers.UploadObject)
}

func RunHTTPServer(server *http.ServeMux, addr string) error {
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
	err := http.ListenAndServe(addr, c.Handler(server))
	log.Fatal(err)
	return nil
}
