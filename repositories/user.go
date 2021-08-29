package repositories

import (
	"database/sql"

	"harvest/core"
	"harvest/entities"
)

type User struct {
	*sql.DB
}

func (r User) Create(u entities.User) error {
	_, err := r.Exec(
		`insert into users (
			firebase_uid, first_name, last_name
		) values (?, ?, ?)`,
		u.Uid, u.FirstName, u.LastName,
	)
	if err != nil {
		switch err {
		case sql.ErrConnDone:
			return core.DSConnErr{Original: err}
		default:
			return core.UnknownErr{Original: err}
		}
	}

	return nil
}
