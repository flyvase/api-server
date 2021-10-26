package controller

import (
	"harvest/domain/entity"
	"harvest/domain/repository"
)

func CreateUser(user entity.User, userR repository.User, authR repository.Auth) error {
	return userR.Create(user)
}
