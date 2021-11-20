package repository

import (
	"harvest/src/domain/model"
	"harvest/src/domain/value/space"
)

type Space interface {
	Fetch(space.Id) (*model.Space, error)
	List() ([]*model.Space, error)
}
