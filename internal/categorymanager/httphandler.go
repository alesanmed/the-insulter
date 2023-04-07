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

type categoryController struct {
	svc *categoryService
}

func NewController(service *categoryService) categoryController {
	return categoryController{
		svc: service,
	}
}

func (controller *categoryController) GetAllCategoriesController(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	categories := controller.svc.GetAllCategories()

	encoder.Encode(categories)

	w.WriteHeader(http.StatusOK)
}

func (controller *categoryController) CreateCategoryController(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var category CreateCategoryDto

	err := decoder.Decode(&category)
	if err != nil {
		log.Printf("error decoding CreateCategoryDto %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := controller.svc.CreateCategory(category.Name)
	if err != nil {
		log.Printf("error creating category %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)

	w.WriteHeader(http.StatusCreated)
	encoder.Encode(CreateCategoryResponse{
		ID: id,
	})
}
