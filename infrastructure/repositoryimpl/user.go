package repositoryimpl

import (
	"harvest/domain/entity"
	"harvest/infrastructure/sql"
)

type User struct {
	Sql sql.Sql
}

func (u *User) Create(user entity.User) (int64, error) {
	res, err := u.Sql.Exec(
		`insert into users (
			firebase_uid, first_name, last_name
		) values (?, ?, ?)`, user.Uid, user.FirstName, user.LastName,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
