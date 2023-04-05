package videomanager

import (
	"github.com/alesanmed/the-insulter/pkg/database"
	"gorm.io/gorm"
)

type VideoRepository interface {
	GetAllVideos() []database.Video
	CreateVideo(name string, extension string, categories []uint) (id uint, err error)
}

type GormVideoRepository struct {
	DB *gorm.DB
}

func (videoRepository *GormVideoRepository) GetAllVideos() (v []database.Video) {
	videoRepository.DB.Find(&v)

	return
}

func (VideoRepository *GormVideoRepository) CreateVideo(name string, extension string, categories []uint) (id uint, err error) {
	categoriesModel := make([]database.Category, len(categories))

	for i := range categories {
		categoriesModel[i] = database.Category{
			Model: gorm.Model{ID: categories[i]},
		}
	}

	video := database.Video{
		Url: name + extension,
		Categories: []database.Category{
			{
				Model: gorm.Model{ID: 1},
			},
		},
	}

	tx := VideoRepository.DB.Omit("Categories.*").Create(&video)

	if err = tx.Error; err != nil {
		return
	}

	id = video.ID

	return
}
