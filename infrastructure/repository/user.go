package repository

import (
	"harvest/domain/entity"
	"harvest/domain/repository"
)

type UserImpl struct {
	Sql repository.Sql
}

func (u UserImpl) Create(user entity.User) (int64, error) {
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
