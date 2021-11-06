package repository

import "harvest/domain/entity"

type Space interface {
	List() ([]*entity.Space, error)
	Fetch(uint32) (*entity.Space, error)
}
