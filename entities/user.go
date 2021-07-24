package entities

import (
	"time"

	_ "github.com/go-playground/validator"
)

type User struct {
	Id uint32 `json:"id" validate:"required"`
	// EXCEPTION: Generally for Firebase
	Uid       string    `json:"uid" validate:"required,lt=255"`
	FirstName string    `json:"first_name" validate:"required,lt=100"`
	LastName  string    `json:"last_name" validate:"required,lt=100"`
	CreatedAt time.Time `json:"created_at,omit_empty"`
	UpdatedAt time.Time `json:"update_at,omit_empty"`
	DeletedAt time.Time `json:"deleted_at,omit_empty"`
}
