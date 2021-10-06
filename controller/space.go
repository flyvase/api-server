package controller

import (
	"harvest/domain/repository"
)

func FetchSpace(spaceR repository.Space) error {
	err := spaceR.Fetch()

	if err != nil {
		return err
	}

	return nil
}
