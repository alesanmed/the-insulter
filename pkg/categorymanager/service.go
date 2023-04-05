package categorymanager

import "github.com/alesanmed/the-insulter/pkg/database"

type CategoryService struct {
	categoryRepository CategoryRepository
}

func (categoryService *CategoryService) GetAllCategories() []database.Category {
	return categoryService.categoryRepository.GetAllCategories()
}

func (categoryService *CategoryService) CreateCategory(name string) (id uint, err error) {
	id, err = categoryService.categoryRepository.CreateCategory(name)

	return
}
