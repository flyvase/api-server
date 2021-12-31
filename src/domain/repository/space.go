package repository

import (
	"api-server/src/domain/model"
	"api-server/src/domain/value"
)

type Space interface {
	List() ([]*model.Space, error)
	Fetch(value.SpaceId) (*model.Space, error)
	GetWebsiteUrl(value.SpaceId) (string, error)
}
