package categorymanager

import (
	"github.com/alesanmed/the-insulter/internal/app"
	"github.com/alesanmed/the-insulter/internal/database"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {
	repository := NewPostgresRepository(database.GetDB())
	service := NewService(&repository)
	controller := NewController(&service)

	r.Get("/category", app.HandlerWithErrors(controller.GetAllCategoriesController))
	r.Post("/category", app.HandlerWithErrors(controller.CreateCategoryController))
}
