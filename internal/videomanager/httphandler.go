package videomanager

import (
	"encoding/json"
	"fmt"
	"log"
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

func (controller *videoController) GetVideosController(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	videos := controller.svc.GetAllVideos()

	encoder.Encode(videos)

	w.WriteHeader(http.StatusOK)
}

func (controller *videoController) CreateVideoController(w http.ResponseWriter, r *http.Request) {
	reader, err := r.MultipartReader()
	if err != nil {
		log.Printf("error creating form reader, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	created_id, err := controller.svc.CreateVideo(reader)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{ "id": %d }`, created_id)))
}
