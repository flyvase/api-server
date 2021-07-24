package interfaces

import "harvest/entities"

type User interface {
	Create(entities.User) error
}
