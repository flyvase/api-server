package entity

import (
	"time"
)

type User struct {
	Id        uint32
	Uid       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
