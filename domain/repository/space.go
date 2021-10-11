package repository

import "harvest/domain/entity"

type Space interface {
	Fetch() ([]entity.Space, error)
}
