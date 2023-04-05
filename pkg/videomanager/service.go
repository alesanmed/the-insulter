package videomanager

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strconv"

	"github.com/alesanmed/the-insulter/pkg/database"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type VideoService struct {
	videoRepository VideoRepository
}

func (videoService *VideoService) GetAllVideos() []database.Video {
	return videoService.videoRepository.GetAllVideos()
}

func (videoService *VideoService) CreateVideo(reader *multipart.Reader) (created_id uint, err error) {
	var video_name string
	categories := make([]uint, 0, 10)

	video_id, err := uuid.NewRandom()
	if err != nil {
		log.Printf("error creating video UUID, %v\n", err)
		return
	}
	video_path := viper.GetString("VIDEO_FOLDER") + video_id.String()

	video_writer, err := os.Create(video_path)
	if err != nil {
		log.Printf("error creating video file, %v\n", err)
		return
	}
	defer video_writer.Close()

	for part, part_err := reader.NextRawPart(); part_err != io.EOF; part, part_err = reader.NextPart() {
		if part.FormName() == "video" {
			if video_name == "" {
				video_name = part.FileName()
			}

			if written, inner_err := io.Copy(video_writer, part); inner_err != nil {
				log.Printf("error creating video file, %v\n", inner_err)
				err = inner_err
				return
			} else {
				log.Printf("written %d bytes\n", written)
			}
		} else if part.FormName() == "category" {
			var category string
			categoryBytes := make([]byte, 512)

			for n, err := part.Read(categoryBytes); ; {
				if err != nil {
					if err == io.EOF {
						category += string(categoryBytes[:n])
						break
					} else {
						category += string(categoryBytes)
					}
				} else {
					break
				}
			}

			parsedCategory, inner_err := strconv.Atoi(category)

			if inner_err != nil {
				log.Printf("error parsing category, %v\n", inner_err)
				err = inner_err
				return
			}

			categories = append(categories, uint(parsedCategory))
		}
	}

	video_ext := path.Ext(video_name)
	os.Rename(video_path, video_path+video_ext)

	created_id, err = videoService.videoRepository.CreateVideo(path.Base(video_name), video_ext, categories)

	return
}
