package repositoryimpl

import (
	"fmt"
	"harvest/core/logger"
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

func (s *Space) Fetch() (error) {
	rows, err := s.Sql.Query(
		"select * from spaces",
	)

	if err != nil {
		return err
	}

	// 以下サンプルをコピペ中
	defer rows.Close()

	var s2 Spaces
	for rows.Next() {
		err := rows.Scan(&s2.ID, &s2.Name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", s2.ID, s2.Name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	logger.Debug("ok", "space")

	return nil
}
