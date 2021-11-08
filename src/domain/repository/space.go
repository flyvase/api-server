package repository

import "harvest/src/domain/entity"

type Space interface {
	List() ([]*entity.Space, error)
	Fetch(uint32) (*entity.Space, error)
}
