package repositoryimpl

import (
	"harvest/domain/entity"
	"harvest/infrastructure/sql"
	"log"
)

type Space struct {
	Sql sql.Sql
}

type Spaces struct {
	ID   uint32 `db:"id"`
	Name string `db:"name"`
}

func (s *Space) Fetch() ([]entity.Space, error) {
	rows, err := s.Sql.Query(
		"select * from spaces",
	)

	var se entity.Space
	var spaces []entity.Space

	if err != nil {
		return spaces, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&se.Id, &se.Name)
		if err != nil {
			log.Fatal(err)
		}

		spaces = append(spaces, se)
	}

	err = rows.Err()

	if err != nil {
		log.Fatal(err)
	}

	return spaces, nil
}
