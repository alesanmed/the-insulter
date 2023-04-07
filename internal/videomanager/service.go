package videomanager

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type videoService struct {
	accepted_video_formats map[string]struct{}
	repository             *VideoRepository
}

func NewService(repository *VideoRepository) videoService {
	return videoService{
		accepted_video_formats: map[string]struct{}{"video/mp4": {}, "video/3gpp": {}},
		repository:             repository,
	}
}

func (svc *videoService) GetAllVideos() []Video {
	videoModels := (*svc.repository).GetAllVideos()

	videos := make([]Video, len(videoModels))

	for i, video := range videoModels {
		videos[i] = Video{
			Id:         video.Id,
			Name:       video.Name,
			Url:        video.Url,
			CreatedAt:  video.CreatedAt,
			UpdatedAt:  video.UpdatedAt,
			Categories: video.Categories,
		}
	}

	return videos
}

func (svc *videoService) CreateVideo(reader *multipart.Reader) (created_id uint, err error) {
	var video_name string
	var video_ext string
	categories := make([]uint, 0, 10)

	video_id, err := uuid.NewRandom()
	if err != nil {
		log.Printf("error creating video UUID, %v", err)
		return
	}
	video_path := viper.GetString("VIDEO_FOLDER") + video_id.String()

	video_writer, err := os.Create(video_path)
	if err != nil {
		log.Printf("error creating video file, %v", err)
		return
	}
	defer func() {
		video_writer.Close()
		if err != nil {
			os.Remove(video_path)
		}
	}()

	for part, part_err := reader.NextRawPart(); part_err != io.EOF; part, part_err = reader.NextPart() {
		if part.FormName() == "video" {
			mime_type, _, err := mime.ParseMediaType(part.Header.Get("Content-Type"))
			if err != nil {
				log.Printf("error parsing file type %v", err)
				return 0, err
			}

			if _, ok := svc.accepted_video_formats[mime_type]; !ok {
				log.Printf("file type not accepted, %s", mime_type)
				err = errors.New("invalid file type")
				return 0, err
			}

			video_ext = filepath.Ext(part.FileName())

			if written, err := io.Copy(video_writer, part); err != nil {
				log.Printf("error creating video file, %v", err)
				return 0, err
			} else {
				log.Printf("written %d bytes", written)
			}
		} else if part.FormName() == "category" {
			var category string
			category_bytes := make([]byte, 512)

			for n, err := part.Read(category_bytes); ; {
				if err != nil {
					if err == io.EOF {
						category += string(category_bytes[:n])
						break
					} else {
						fmt.Printf("err parsing category %v", err)
						return 0, err
					}
				} else {
					category += string(category_bytes)
				}
			}

			parsed_category, err := strconv.Atoi(category)

			if err != nil {
				log.Printf("error parsing category, %v", err)
				return 0, err
			}

			categories = append(categories, uint(parsed_category))
		} else if part.FormName() == "name" {
			video_name_bytes := make([]byte, 512)

			for n, err := part.Read(video_name_bytes); ; {
				if err != nil {
					if err == io.EOF {
						video_name += string(video_name_bytes[:n])
						break
					} else {
						fmt.Printf("err parsing video name %v", err)
						return 0, err
					}
				} else {
					video_name += string(video_name_bytes)
				}
			}
		}
	}

	old_path, video_path := video_path, video_path+video_ext
	os.Rename(old_path, video_path)

	created_id, err = (*svc.repository).CreateVideo(video_name, video_path, categories)

	return
}
