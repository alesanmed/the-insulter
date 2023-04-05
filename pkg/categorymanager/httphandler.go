package categorymanager

import (
	"encoding/json"
	"log"
	"net/http"
)

type CreateCategoryDto struct {
	Name string `json:"name"`
}

type CreateCategoryResponse struct {
	ID uint `json:"id"`
}

type CategoryController struct {
	categoryService *CategoryService
}

func (categoryController *CategoryController) GetAllCategoriesController(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	categories := categoryController.categoryService.GetAllCategories()

	encoder.Encode(categories)

	w.WriteHeader(http.StatusOK)
}

func (categoryController *CategoryController) CreateCategoryController(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var category CreateCategoryDto

	err := decoder.Decode(&category)
	if err != nil {
		log.Printf("error decoding CreateCategoryDto %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	id, err := categoryController.categoryService.CreateCategory(category.Name)
	if err != nil {
		log.Printf("error creating category %v\n", err)
	}

	encoder := json.NewEncoder(w)

	w.WriteHeader(http.StatusCreated)
	encoder.Encode(CreateCategoryResponse{
		ID: id,
	})
}
