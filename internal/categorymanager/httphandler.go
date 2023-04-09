package categorymanager

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alesanmed/the-insulter/internal/app"
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

func (controller *categoryController) GetAllCategoriesController(w http.ResponseWriter, r *http.Request) (err error) {
	encoder := json.NewEncoder(w)

	categories := controller.svc.GetAllCategories()

	encoder.Encode(categories)

	w.WriteHeader(http.StatusOK)

	return
}

func (controller *categoryController) CreateCategoryController(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)

	var category CreateCategoryDto

	err := decoder.Decode(&category)
	if err != nil {
		return fmt.Errorf("error decoding CreateCategoryDto: %w", app.NewAPIError(http.StatusBadRequest, "Invalid category body", err))
	}

	id, err := controller.svc.CreateCategory(category.Name)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(w)

	w.WriteHeader(http.StatusCreated)
	encoder.Encode(CreateCategoryResponse{
		ID: id,
	})

	return nil
}
