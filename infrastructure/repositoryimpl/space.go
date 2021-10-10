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
		err := rows.Scan(&se.Id, &se.Name)
		if err != nil {
			return nil, err
		}

		spaces = append(spaces, se)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return spaces, nil
}
