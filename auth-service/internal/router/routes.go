package router

import (
	"log"
	"net/http"

	"github.com/DanialKassym/GoStorage/auth-service/internal/handlers"
	"github.com/rs/cors"
)

func InitRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register/", handlers.Register)
	mux.HandleFunc("POST /login/", handlers.Login)
	mux.HandleFunc("GET /main/", handlers.Main)
	mux.HandleFunc("POST /validate/", handlers.CheckJWT)
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
