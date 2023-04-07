package videomanager

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/lib/pq"
)

type VideoRepository interface {
	GetAllVideos() []videoModel
	CreateVideo(name string, path string, categories []uint) (id uint, err error)
}

type postgresVideoRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(db *sql.DB) VideoRepository {
	return postgresVideoRepository{
		DB: db,
	}
}

func (repository postgresVideoRepository) GetAllVideos() (v []videoModel) {
	rows, err := repository.DB.Query("select v.id, v.name, v.url, v.created_at, v.updated_at, array_agg(json_build_object('id', vc.category_id, 'name', c.name)) as categories from videos v inner join video_categories vc on v.id = vc.video_id inner join categories c on vc.category_id = c.id where v.deleted_at is null and c.deleted_at is null group by v.id")
	if err != nil {
		log.Printf("error getting videos %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id            uint
			name          string
			url           string
			created_at    time.Time
			updated_at    time.Time
			rawCategories []string
		)

		rows.Scan(&id, &name, &url, &created_at, &updated_at, pq.Array(&rawCategories))

		categories := make([]videoCategory, len(rawCategories))

		for i, rawCategory := range rawCategories {
			err := json.Unmarshal([]byte(rawCategory), &categories[i])
			if err != nil {
				log.Printf("error processing video categories %v", err)
				return
			}
		}

		v = append(v, videoModel{
			Id:         id,
			Name:       name,
			Url:        url,
			CreatedAt:  created_at,
			UpdatedAt:  updated_at,
			DeletedAt:  sql.NullTime{},
			Categories: categories,
		})
	}

	return
}

func (repository postgresVideoRepository) CreateVideo(name string, path string, categories []uint) (id uint, err error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		log.Printf("error starting video transaction %v", err)
		return
	}
	defer tx.Rollback()

	err = tx.QueryRow("insert into videos (id, name, url) values (default, $1, $2) returning id", name, path).Scan(&id)
	if err != nil {
		log.Printf("error inserting video %v", err)
		return
	}

	for _, category_id := range categories {
		_, err = tx.Exec("insert into video_categories (video_id, category_id) values ($1, $2)", id, category_id)
		if err != nil {
			log.Printf("error inserting category association %v", err)
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.Printf("error committing video transaction %v", err)
		return
	}

	return
}
