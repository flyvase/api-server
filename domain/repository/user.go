package repository

import "harvest/domain/entity"

type User interface {
	Create(entity.User) (int64, error)
}