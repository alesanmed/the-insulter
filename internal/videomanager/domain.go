package videomanager

import "time"

type Video struct {
	Id         uint
	Name       string
	Url        string
	Categories []videoCategory
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CustomizedVideo struct {
	Id        uint
	Name      string
	Url       string
	UserId    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
