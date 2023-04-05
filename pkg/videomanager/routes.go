package videomanager

import (
	"github.com/alesanmed/the-insulter/pkg/database"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {
	videoController := &VideoController{
		videoService: &VideoService{
			videoRepository: &GormVideoRepository{
				DB: database.GetDB(),
			},
		},
	}

	r.Get("/video", videoController.GetVideosController)
	r.Post("/video", videoController.CreateVideoController)
}
