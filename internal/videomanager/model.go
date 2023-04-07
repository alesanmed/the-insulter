package videomanager

import (
	"database/sql"
	"time"
)

type videoCategory struct {
	Id   uint
	Name string
}

type videoModel struct {
	Id         uint
	Name       string
	Url        string
	Categories []videoCategory
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime
}

type customizedVideoModel struct {
	Id        uint
	Name      string
	Url       string
	UserId    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
