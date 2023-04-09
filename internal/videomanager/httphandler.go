package videomanager

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type videoController struct {
	svc *videoService
}

func NewController(svc *videoService) videoController {
	return videoController{
		svc: svc,
	}
}

func (controller *videoController) GetVideosController(w http.ResponseWriter, r *http.Request) error {
	encoder := json.NewEncoder(w)

	videos := controller.svc.GetAllVideos()

	encoder.Encode(videos)

	w.WriteHeader(http.StatusOK)

	return nil
}

func (controller *videoController) CreateVideoController(w http.ResponseWriter, r *http.Request) (err error) {
	reader, err := r.MultipartReader()
	if err != nil {
		return err
	}

	created_id, err := controller.svc.CreateVideo(reader)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{ "id": %d }`, created_id)))

	return
}
