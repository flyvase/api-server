package repository

import (
	"harvest/src/domain/entity"
	"harvest/src/infrastructure/sql"
)

type User struct {
	Driver sql.Driver
}

func (ur *User) Create(user entity.User) error {
	err := ur.Driver.Exec(
		`insert into users (
			firebase_uid
		) values (?)`, user.Uid,
	)
	return err
}
