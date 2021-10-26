package controller

import (
	"harvest/domain/entity"
	"harvest/domain/repository"
)

func ListSpace(spaceR repository.Space) ([]entity.Space, error) {
	return spaceR.List()
}
