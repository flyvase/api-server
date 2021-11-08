package repository

import "harvest/src/domain/entity"

type User interface {
	Create(entity.User) error
}
