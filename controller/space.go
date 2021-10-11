package controller

import (
	"harvest/domain/entity"
	"harvest/domain/repository"
)

func FetchSpace(spaceR repository.Space) ([]entity.Space, error) {
	spaces, err := spaceR.Fetch()

	if err != nil {
		return nil, err
	}

	return spaces, nil
}
