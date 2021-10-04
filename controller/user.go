package controller

import (
	"strconv"

	"harvest/domain/entity"
	"harvest/domain/repository"
)

func CreateUser(user entity.User, userR repository.User, authR repository.Auth) error {
	id, err := userR.Create(user)
	if err != nil {
		return err
	}

	claims := map[string]interface{}{"id": strconv.FormatInt(id, 10)}
	err = authR.SetCustomClaim(user, claims)
	if err != nil {
		return err
	}

	return nil
}
