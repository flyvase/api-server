package repository

import (
	"harvest/domain/entity"
	"harvest/infrastructure/sql"
)

type User struct {
	Driver sql.Driver
}

func (ur *User) Create(user entity.User) error {
	err := ur.Driver.Exec(
		`insert into users (
			firebase_uid, first_name, last_name
		) values (?, ?, ?)`, user.Uid, user.FirstName, user.LastName,
	)
	return err
}