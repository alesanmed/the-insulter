package main

import (
	"log"
	"net/http"

	"github.com/alesanmed/the-insulter/pkg/config"
	"github.com/alesanmed/the-insulter/pkg/database"
	"github.com/alesanmed/the-insulter/pkg/routes"
	"github.com/go-chi/chi/v5"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("cannot read config %v", err)
	}

	r := chi.NewRouter()

	database.Init()

	routes.RegisterRoutes(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Cannot open server %v", err)
	}

	log.Println("Server up and running")
}
