package controller

import (
	"harvest/domain/entity"
	"harvest/domain/repository"
)

func CreateUser(u entity.User, r repository.User) error {
	_, err := r.Create(u)
	if err != nil {
		return err
	}

	// TODO: register returned id to firebase authentication custom claim

	return nil
}
