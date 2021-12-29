package repository

import (
	"harvest/src/domain/model"
	"harvest/src/domain/value"
)

type Space interface {
	List() ([]*model.Space, error)
	Fetch(value.SpaceId) (*model.Space, error)
	GetWebsiteUrl(value.SpaceId) (string, error)
}
