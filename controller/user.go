package controller

import (
	"strconv"

	"harvest/domain/entity"
	"harvest/domain/repository"
)

func CreateUser(u entity.User, r repository.User, a repository.Auth) error {
	id, err := r.Create(u)
	if err != nil {
		return err
	}

	// TODO: register returned id to firebase authentication custom claim
	err = a.SetCustomClaim(strconv.FormatInt(id, 10))
	if err != nil {
		return err
	}

	return nil
}
