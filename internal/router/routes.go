package routes

import (
	"log"
	"net/http"

	"github.com/DanialKassym/GoStorage/internal/handlers"
)

func InitRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.RetriveUsers)
	mux.HandleFunc("POST /upload/", handlers.UploadObject)
	log.Println("starting server on 8080")
	err := http.ListenAndServe("0.0.0.0:8080", mux)
	log.Fatal(err)
}
