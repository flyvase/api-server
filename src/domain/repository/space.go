package repository

import (
	"harvest/src/domain/model"
	"harvest/src/domain/value"
)

type Space interface {
	Fetch(value.SpaceId) (*model.Space, error)
	List() ([]*model.Space, error)
}
