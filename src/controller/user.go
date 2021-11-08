package controller

import (
	"harvest/src/domain/entity"
	"harvest/src/domain/repository"
)

func CreateUser(user entity.User, userR repository.User, authR repository.Auth) error {
	return userR.Create(user)
}
