package router

import (
	"log"
	"net/http"

	"github.com/DanialKassym/GoStorage/cmd/authentication/internal/handlers"
	"github.com/rs/cors"
)

func InitRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /login/", handlers.Authorize)
	mux.HandleFunc("POST /register/", handlers.Register)
	mux.HandleFunc("GET /get/", handlers.Get)
	log.Println("starting server on 8081")
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
	err := http.ListenAndServe("0.0.0.0:8081", c.Handler(mux))
	log.Fatal(err)
}
