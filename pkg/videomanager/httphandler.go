package videomanager

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type VideoController struct {
	videoService *VideoService
}

func (videoController *VideoController) GetVideosController(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	videos := videoController.videoService.GetAllVideos()

	encoder.Encode(videos)

	w.WriteHeader(http.StatusOK)
}

func (videoController *VideoController) CreateVideoController(w http.ResponseWriter, r *http.Request) {
	reader, err := r.MultipartReader()
	if err != nil {
		log.Printf("error creating form reader, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	created_id, err := videoController.videoService.CreateVideo(reader)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("{ 'video_id': %d }", created_id)))
}
