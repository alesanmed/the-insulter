package categorymanager

type categoryService struct {
	repository *CategoryRepository
}

func NewService(repository *CategoryRepository) categoryService {
	return categoryService{
		repository: repository,
	}
}

func (svc *categoryService) GetAllCategories() []Category {
	categories := (*svc.repository).GetAllCategories()

	res := make([]Category, len(categories))

	for i, category := range categories {
		res[i] = Category{
			Id:        category.Id,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		}
	}

	return res
}

func (svc *categoryService) CreateCategory(name string) (id uint, err error) {
	id, err = (*svc.repository).CreateCategory(name)

	return
}
