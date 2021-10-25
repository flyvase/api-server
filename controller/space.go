package controller

import (
	"harvest/domain/entity"
	"harvest/domain/repository"
)

func ListSpace(spaceR repository.Space) ([]entity.Space, error) {
	spaces, err := spaceR.List()

	if err != nil {
		return nil, err
	}

	return spaces, nil
}
