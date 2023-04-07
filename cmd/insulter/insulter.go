package main

import (
	"log"
	"net/http"

	"github.com/alesanmed/the-insulter/internal/config"
	"github.com/alesanmed/the-insulter/internal/database"
	"github.com/alesanmed/the-insulter/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("cannot read config %v", err)
	}

	r := chi.NewRouter()

	database.Init()

	db := database.GetDB()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("cannot create postgres driver %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations/", "postgres", driver)
	if err != nil {
		log.Fatalf("cannot create migrations instance %v", err)
	}
	if err = m.Up(); err != nil && err.Error() != "no change" {
		log.Fatalf("error running migrations %v", err)
	}

	routes.RegisterRoutes(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Cannot open server %v", err)
	}

	log.Println("Server up and running")
}
