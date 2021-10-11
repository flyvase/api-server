package repositoryimpl

import (
	"harvest/domain/entity"
	"harvest/infrastructure/sql"
)

type Space struct {
	Sql sql.Sql
}

func (s *Space) Fetch() ([]entity.Space, error) {
	rows, err := s.Sql.Query(
		"select * from spaces",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var spaces []entity.Space

	for rows.Next() {
		var se entity.Space
		if err := rows.Scan(&se.Id, &se.Name); err != nil {
			return nil, err
		}

		spaces = append(spaces, se)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return spaces, nil
}
