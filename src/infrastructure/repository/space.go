package repository

import (
	"harvest/src/domain/entity"
	"harvest/src/infrastructure/sql"
)

type Space struct {
	Driver sql.Driver
}

func (sr *Space) List() ([]*entity.Space, error) {
	rows, err := sr.Driver.Query(
		"select * from spaces",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var spaces []*entity.Space

	for rows.Next() {
		var se entity.Space
		if err := rows.Scan(&se.Id, &se.Name); err != nil {
			return nil, err
		}

		spaces = append(spaces, &se)
	}

	return spaces, nil
}

func (sr *Space) Fetch(id uint32) (*entity.Space, error) {
	row, err := sr.Driver.QueryRow(
		`select * from spaces where id = ?`,
		id,
	)

	if err != nil {
		return nil, err
	}

	var se entity.Space
	if err := row.Scan(&se.Id, &se.Name); err != nil {
		return nil, err
	}

	return &se, nil
}
