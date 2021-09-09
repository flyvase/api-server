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

	err = a.SetCustomClaim(u, strconv.FormatInt(id, 10))
	if err != nil {
		return err
	}

	return nil
}
