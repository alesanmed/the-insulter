package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/alesanmed/the-insulter/internal/categorymanager"
	"github.com/alesanmed/the-insulter/internal/middlewares"
	"github.com/alesanmed/the-insulter/internal/videomanager"
)

func RegisterRoutes(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json", "multipart/form-data"))
	r.Use(middlewares.ResponseHeaderMiddleware)

	videomanager.RegisterRoutes(r)
	categorymanager.RegisterRoutes(r)
}
