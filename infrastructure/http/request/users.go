package request

import (
	"encoding/json"
	"harvest/domain/entity"
	"io"

	"github.com/go-playground/validator"
)

type User struct {
	Uid string `json:"uid" validate:"required,max=255"`
}

func (u User) ToUserEntity() entity.User {
	return entity.User{
		Uid: u.Uid,
	}
}

func DecodeUsersPostRequest(r io.Reader, u *User) error {
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
