package model

import "api-server/src/domain/value"

type SpaceDisplay struct {
	Id          value.SpaceDisplayId
	ImageUrl    string
	Description string
}
