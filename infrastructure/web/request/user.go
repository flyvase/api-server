package request

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"

	"harvest/domain/entity"
)

type User struct {
	Uid       string `json:"uid" validate:"required,max=255"`
	FirstName string `json:"first_name" validate:"required,max=100"`
	LastName  string `json:"last_name" validate:"required,max=100"`
}

func (u User) ToUserEntity() entity.User {
	return entity.User{
		Uid:       u.Uid,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

func DecodeUserRequestJson(r io.Reader, u *User) error {
	dec := json.NewDecoder(r)
	if err := dec.Decode(u); err != nil {
		return err
	}

	v := validator.New()
	if err := v.Struct(u); err != nil {
		return err
	}

	return nil
}
