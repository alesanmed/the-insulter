package categorymanager

import (
	"database/sql"
	"time"
)

type categoryModel struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
