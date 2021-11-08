package controller

import (
	"harvest/src/domain/entity"
	"harvest/src/domain/repository"
)

func ListSpaces(spaceR repository.Space) ([]*entity.Space, error) {
	return spaceR.List()
}

func FetchSpace(id uint32, spaceR repository.Space) (*entity.Space, error) {
	return spaceR.Fetch(id)
}
