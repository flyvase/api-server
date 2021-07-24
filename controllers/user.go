package controllers

import (
	"harvest/entities"
	"harvest/interfaces"
)

func CreateUser(i interfaces.User, u entities.User) error {
	if err := i.Create(u); err != nil {
		return err
	}

	return nil
}
