package entities

import (
	"time"

	_ "github.com/go-playground/validator"
)

type User struct {
	Id uint32 `json:"id,omit_empty"`
	// EXCEPTION: Generally for Firebase
	Uid       string    `json:"uid" validate:"required,max=255"`
	FirstName string    `json:"first_name" validate:"required,max=100"`
	LastName  string    `json:"last_name" validate:"required,max=100"`
	CreatedAt time.Time `json:"created_at,omit_empty"`
	UpdatedAt time.Time `json:"update_at,omit_empty"`
	DeletedAt time.Time `json:"deleted_at,omit_empty"`
}
