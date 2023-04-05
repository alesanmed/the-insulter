package categorymanager

import (
	"github.com/alesanmed/the-insulter/pkg/database"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategories() []database.Category
	CreateCategory(name string) (id uint, err error)
}

type GormCategoryRepository struct {
	DB *gorm.DB
}

func (categoryRepository *GormCategoryRepository) GetAllCategories() (categories []database.Category) {
	categoryRepository.DB.Find(&categories)

	return
}

func (categoryRepository *GormCategoryRepository) CreateCategory(name string) (id uint, err error) {
	category := database.Category{
		Name: name,
	}

	tx := categoryRepository.DB.Create(&category)

	if err = tx.Error; err != nil {
		return
	}

	id = category.ID

	return
}
