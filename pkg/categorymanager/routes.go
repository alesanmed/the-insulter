package categorymanager

import (
	"github.com/alesanmed/the-insulter/pkg/database"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {
	categoryController := &CategoryController{
		categoryService: &CategoryService{
			categoryRepository: &GormCategoryRepository{
				DB: database.GetDB(),
			},
		},
	}

	r.Get("/category", categoryController.GetAllCategoriesController)
	r.Post("/category", categoryController.CreateCategoryController)
}
