package categorymanager

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/alesanmed/the-insulter/internal/app"
)

type CategoryRepository interface {
	GetAllCategories() []categoryModel
	CreateCategory(name string) (id uint, err error)
}

type postgresCategoryRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(db *sql.DB) CategoryRepository {
	return postgresCategoryRepository{
		DB: db,
	}
}

func (repository postgresCategoryRepository) GetAllCategories() (categories []categoryModel) {
	rows, err := repository.DB.Query("select id, name, created_at, updated_at from categories where deleted_at is null")
	if err != nil {
		log.Printf("error querying categories %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id        uint
			name      string
			createdAt time.Time
			updatedAt time.Time
		)

		rows.Scan(&id, &name, &createdAt, &updatedAt)

		categories = append(categories, categoryModel{
			Id:        id,
			Name:      name,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: sql.NullTime{},
		})
	}

	return
}

func (repository postgresCategoryRepository) CreateCategory(name string) (id uint, err error) {
	sql := "insert into categories (id, name) values (default, $1) RETURNING id"
	if err = repository.DB.QueryRow(sql, name).Scan(&id); err != nil {
		return 0, fmt.Errorf("error inserting category: %w", app.NewAPIError(app.ErrInternal.GetStatus(), app.ErrInternal.GetMessage(), err))
	}

	return
}
