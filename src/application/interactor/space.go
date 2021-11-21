package interactor

import (
	"harvest/src/domain/model"
	"harvest/src/domain/repository"
	"harvest/src/domain/value"
)

type Space struct {
	SpaceRepository repository.Space
}

func (s *Space) List() ([]*model.Space, error) {
	return s.SpaceRepository.List()
}

func (s *Space) Fetch(id value.SpaceId) (*model.Space, error) {
	return s.SpaceRepository.Fetch(id)
}
